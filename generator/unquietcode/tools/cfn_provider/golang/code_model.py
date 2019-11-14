# from enum import Enum, auto, unique
# from collections import defaultdict
from typing import List
from dataclasses import dataclass, field #, asdict, field


@dataclass(frozen=True)
class GoParameter:
    name: str
    type: str
    varargs: bool = False


@dataclass(frozen=True)
class GoFunction:
    name: str
    body: str
    parameters: List[GoParameter] = field(default_factory=lambda: [])
    return_types: List[str] = field(default_factory=lambda: [])
    
    def write(self, writer):
        writer.write_function(self)
        

# https://golang.org/ref/spec#Source_file_organization
@dataclass(frozen=True)
class SourceFile:
    header: str
    package_name: str
    imports: List[str]
    declarations: List[GoFunction]
    
    def write(self, writer):
        writer.write_file(self)






@dataclass(frozen=True)
class GoPackage:
    files: List[SourceFile]





