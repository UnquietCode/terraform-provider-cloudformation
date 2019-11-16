# from enum import Enum, auto, unique
# from collections import defaultdict
from typing import List
from dataclasses import dataclass, field #, asdict, field


@dataclass(frozen=True)
class CodeElement:
    pass


@dataclass(frozen=True)
class BlankLines(CodeElement):
    count: int = 1

    def write(self, writer):
        writer.write_blank_lines(self)


@dataclass(frozen=True)
class AssignmentExpression(CodeElement):
    identifier: str
    type: str
    expression: any

    def write(self, writer):
        writer.write_assignment_expression(self)


@dataclass(frozen=True)
class ConstantExpression(AssignmentExpression):
    def write(self, writer):
        writer.write_constant_expression(self)


@dataclass(frozen=True)
class ReturnExpression(CodeElement):
    expression: any

    def write(self, writer):
        writer.write_return_expression(self)


@dataclass(frozen=True)
class LiteralExpression(CodeElement):
    expression: str

    def write(self, writer):
        writer.write_literal(self)


@dataclass(frozen=True)
class GoParameter:
    name: str
    type: str
    varargs: bool = False


@dataclass(frozen=True)
class GoFunction(CodeElement):
    name: str
    body: List[str] = field(default_factory=lambda: [])
    parameters: List[GoParameter] = field(default_factory=lambda: [])
    return_types: List[str] = field(default_factory=lambda: [])
    
    def write(self, writer):
        writer.write_function(self)


# https://golang.org/ref/spec#Composite_literals
@dataclass(frozen=True)
class GoStructLiteral(CodeElement):
    type: str
    fields: dict = field(default_factory=lambda: {})
    aligned: bool = False
    
    def write(self, writer):
        writer.write_struct_literal(self)


# https://golang.org/ref/spec#Composite_literals
@dataclass(frozen=True)
class GoMapLiteral(CodeElement):
    key_type: str
    value_type: str
    fields: dict = field(default_factory=lambda: {})
    
    def write(self, writer):
        writer.write_map_literal(self)



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





