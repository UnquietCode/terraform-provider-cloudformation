import os
import json

from unquietcode.tools.cfn_provider.parser import handle_spec
from unquietcode.tools.cfn_provider.renderer import render_provider

THIS_DIR = os.path.dirname(os.path.realpath(__file__))
SPEC_FILE = f"{THIS_DIR}/CloudFormationResourceSpecification-us-east-1.json"
OUT_DIR = f"{THIS_DIR}/go"



def main():
    
    with open(SPEC_FILE, 'r+') as file_:
        data = json.load(file_)
    
    top_level_package = handle_spec(data)
    render_provider(top_level_package, OUT_DIR)


if __name__ == '__main__':
    main()