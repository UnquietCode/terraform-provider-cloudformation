# terraform-provider-cloudformation
Experimental AWS CloudFormation provider for Terraform. Allows for building CloudFormation templates from within Terraform,
leveraging the rich modularization and interpolation features that TF has to offer.

### Experimental!
While presenting an interesting mix of code generation and terraform tooling, this is an unfinished codebase with some
serious challenges to work out.

### What does it look like?
You can take something like this:

```yaml
Resources:
  S3BucketW3:
    Type: "AWS::S3::Bucket"
    Properties: 
      BucketName: !Sub www.${DomainName}
      AccessControl: PublicRead
      Tags:
        - Key: Client
          Value: !Ref Client
      VersioningConfiguration:
        Status: Enabled
      WebsiteConfiguration:
        IndexDocument: index.html
        ErrorDocument: 404.html
  S3BucketBare:
    Type: "AWS::S3::Bucket"
    Properties: 
      BucketName: !Sub ${DomainName}
      AccessControl: PublicRead
      Tags:
        - Key: Client
          Value: !Ref Client
      WebsiteConfiguration:
        RedirectAllRequestsTo:
          HostName: !Sub www.${DomainName}
          Protocol: https
  S3BucketPolicy:
    Type: "AWS::S3::BucketPolicy"
    Properties: 
      Bucket: !Ref S3BucketW3
      PolicyDocument: 
        Version: "2012-10-17"
        Statement:
          -
            Sid: PublicReadGetObject
            Effect: Allow
            Principal: "*"
            Action:
              - s3:GetObject
            Resource: !Sub "arn:aws:s3:::www.${DomainName}/*"
  LogBucket:
    Type: "AWS::S3::Bucket"
    Properties: 
      BucketName: !Sub ${DomainName}-logs
      Tags:
        - Key: Client
          Value: !Ref Client
  SSLCert:
    Type: "AWS::CertificateManager::Certificate"
    Properties: 
      DomainName: !Sub "*.${DomainName}"
      SubjectAlternativeNames:
        - !Ref DomainName
      Tags:
        - Key: Client
          Value: !Ref Client
  CloudFront:
    Type: AWS::CloudFront::Distribution
    Properties:
      DistributionConfig:
        Origins:
        - DomainName: !Sub 
            - ${S3BucketW3}.${S3WebEndpoint}
            - { S3BucketW3: !Ref S3BucketW3, S3WebEndpoint: !FindInMap [RegionMap, !Ref "AWS::Region", websiteendpoint] }
          Id: myS3Origin
          CustomOriginConfig:
            OriginProtocolPolicy: http-only
        Enabled: 'true'
        Comment: !Sub Distribution for ${Client}
        HttpVersion: http2
        Logging:
          IncludeCookies: 'false'
          Bucket: !Sub ${LogBucket}.s3.amazonaws.com
          Prefix: !Ref DomainName
        Aliases:
        - !Sub www.${DomainName}
        DefaultCacheBehavior:
          AllowedMethods:
          - DELETE
          - GET
          - HEAD
          - OPTIONS
          - PATCH
          - POST
          - PUT
          TargetOriginId: myS3Origin
          Compress: True
          MinTTL: 300
          DefaultTTL: 300
          ForwardedValues:
            QueryString: 'false'
            Cookies:
              Forward: none
          ViewerProtocolPolicy: redirect-to-https
        PriceClass: PriceClass_100
        ViewerCertificate:
          AcmCertificateArn: !Ref SSLCert
          SslSupportMethod: sni-only
  CloudFrontBare:
    Type: AWS::CloudFront::Distribution
    Properties:
      DistributionConfig:
        Origins:
        - DomainName: !Sub 
            - ${S3BucketBare}.${S3WebEndpoint}
            - { S3BucketBare: !Ref S3BucketBare, S3WebEndpoint: !FindInMap [RegionMap, !Ref "AWS::Region", websiteendpoint] }
          Id: myS3OriginBare
          CustomOriginConfig:
            OriginProtocolPolicy: http-only
        Enabled: 'true'
        Comment: !Sub Distribution for ${Client}
        HttpVersion: http2
        Logging:
          IncludeCookies: 'false'
          Bucket: !Sub ${LogBucket}.s3.amazonaws.com
          Prefix: !Sub ${DomainName}_Bare
        Aliases:
        - !Ref DomainName
        DefaultCacheBehavior:
          AllowedMethods:
          - DELETE
          - GET
          - HEAD
          - OPTIONS
          - PATCH
          - POST
          - PUT
          TargetOriginId: myS3OriginBare
          Compress: True
          DefaultTTL: 604800
          ForwardedValues:
            QueryString: 'false'
            Cookies:
              Forward: none
          ViewerProtocolPolicy: redirect-to-https
        PriceClass: PriceClass_100
        ViewerCertificate:
          AcmCertificateArn: !Ref SSLCert
          SslSupportMethod: sni-only
```

and instead use something like this:

```hcl
provider "cfn" {
  workdir = "/home/ben/Documents/Code-Projects/cfnpvd/s3_test/workdir"
}

variable domain_name {
  default = "test.unquietcode.net"
}

resource "cfn_s3_bucket" "www_bucket" {
  logical_id = "S3BucketWWW"
  bucket_name = "www.${var.domain_name}"
  access_control = "PublicRead"
  
  versioning_configuration {
    status = "Enabled"
  }
	
  website_configuration {
    index_document = "index.html"
    error_document = "404.html"
  }
}

resource "cfn_s3_bucket" "bare_bucket" {
  logical_id = "S3BucketBare"
  bucket_name = var.domain_name
  access_control = "PublicRead"
	
  website_configuration {
    redirect_all_requests_to {
      host_name = "www.${var.domain_name}"
      protocol = "https"
    }
  }
}

resource "cfn_s3_bucket_policy" "bucket_policy" {
  logical_id = "S3BucketPolicy"
  bucket = cfn_s3_bucket.www_bucket.bucket_name
  
  policy_document = <<EOF
{
  "Version": "2012-10-17",
  "Statement": {
    "Sid": "PublicReadGetObject",
    "Effect": "allow",
    "Principal" :"*",
    "Action": ["s3:GetObject"],
    "Resource": "arn:aws:s3:::www.${var.domain_name}/*"
  }
}
EOF
}

resource "cfn_s3_bucket" "log_bucket" {
  logical_id = "LogBucket"
  bucket_name = "${var.domain_name}-logs"
}

resource "cfn_certificatemanager_certificate" "ssl" {
  logical_id = "SSLCert"
  domain_name = "*.${var.domain_name}"
  subject_alternative_names = [var.domain_name] 
}

resource "cfn_cloudfront_distribution" "cloudfront" {
  logical_id = "CloudFront"
  
  distribution_config {
    origins {
      id = "myS3Origin"
      domain_name = "${cfn_s3_bucket.www_bucket.bucket_name}.s3-website-us-west-2.amazonaws.com"
      custom_origin_config {
        origin_protocol_policy = "http-only"
      }
    }
    enabled = true
    comment = "distribution for website"
    http_version = "http2"
    
    logging {
      include_cookies = false
      bucket = "${cfn_s3_bucket.log_bucket.bucket_name}.s3.amazonaws.com"
      prefix = var.domain_name
    }
    aliases = [
      "www.${var.domain_name}"
    ]
    
    default_cache_behavior {
      allowed_methods = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
      target_origin_id = "myS3Origin"
      compress = true
      min_ttl = 300
      default_ttl = 300
      
      forwarded_values {
        query_string = false
        cookies {
          forward = "none"
        }
      }
      viewer_protocol_policy = "redirect-to-https"
    }
    price_class = "PriceClass_1000"
    viewer_certificate {
      acm_certificate_arn = "!GetAtt ${cfn_certificatemanager_certificate.ssl.logical_id}.arn"
      ssl_support_method = "sni-only"
    }
  }
}


resource "cfn_cloudfront_distribution" "cloudfront_bare" {
  logical_id = "CloudFrontBare"
  
  distribution_config {
    origins {
      id = "myS3OriginBare"
      domain_name = "${cfn_s3_bucket.bare_bucket.bucket_name}.s3-website-us-west-2.amazonaws.com"
      custom_origin_config {
        origin_protocol_policy = "http-only"
      }
    }
    enabled = true
    comment = "distribution for website"
    http_version = "http2"
    
    logging {
      include_cookies = false
      bucket = "${cfn_s3_bucket.log_bucket.bucket_name}.s3.amazonaws.com"
      prefix = "${var.domain_name}_Bare"
    }
    aliases = [
      var.domain_name
    ]
    
    default_cache_behavior {
      allowed_methods = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
      target_origin_id = "myS3OriginBare"
      compress = true
      default_ttl = 604800
      
      forwarded_values {
        query_string = false
        cookies {
          forward = "none"
        }
      }
      viewer_protocol_policy = "redirect-to-https"
    }
    price_class = "PriceClass_200"
    viewer_certificate {
      acm_certificate_arn = "!Ref ${cfn_certificatemanager_certificate.ssl.logical_id}"
      ssl_support_method = "sni-only"
    }
  }
}


resource "cfn_template_data" "stack_template" {
  description = "cool stack yo"
}


resource "local_file" "foo" {
  sensitive_content = cfn_template_data.stack_template.output
  filename = "${path.module}/stack.json"
}
```
