import re


def snake_caps(name):

    # split by capital letters and add underscore
    pretty_name = name

    # before a capital letter not preceded by a capital letter
    pretty_name = re.sub(
        string=pretty_name,
        pattern=r'([^A-Z])([A-Z])',
        repl='\\1_\\2',
    )

    # before a capital letter preceded by a capital letter but followed by lowercase
    pretty_name = re.sub(
        string=pretty_name,
        pattern=r'([A-Z])([A-Z])([^A-Z])',
        repl='\\1_\\2\\3',
    )

    return pretty_name.lower()
