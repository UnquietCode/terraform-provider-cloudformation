import re


def snake_caps(name):
    """
    split by capital letters and add underscore
    """
    pretty_name = name
    
    # special handling for 
    # 'IoT' and 'FSx' and others
    pretty_name = re.sub(r'IoT', 'Iot', pretty_name)
    pretty_name = re.sub(r'FSx', 'Fsx', pretty_name)
    
    # before a capital letter not preceded by a capital letter
    pretty_name = re.sub(
        string=pretty_name,
        pattern=r'([^A-Z])([A-Z])',
        repl='\\1_\\2',
    )

    # before a capital letter preceded by a capital letter but
    # followed by lowercase or number
    pretty_name = re.sub(
        string=pretty_name,
        pattern=r'([A-Z])([A-Z])([^A-Z0-9])',
        repl='\\1_\\2\\3',
    )
    
    return pretty_name.lower()
