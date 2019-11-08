import os
import json

from unquietcode.tools.cfn_provider.parser import handle_provider
from unquietcode.tools.cfn_provider.renderer import render_provider

THIS_DIR = os.path.dirname(os.path.realpath(__file__))
SPEC_FILE = f"{THIS_DIR}/CloudFormationResourceSpecification-us-east-1.json"
OUT_DIR = os.path.abspath(f"{THIS_DIR}/../provider")


def main():
    
    with open(SPEC_FILE, 'r+') as file_:
        data = json.load(file_)
    
    provider = handle_provider(data)
    render_provider(provider, OUT_DIR)


if __name__ == '__main__':
    main()