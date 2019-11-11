from io import StringIO
from dataclasses import dataclass

from unquietcode.tools.cfn_provider.golang.code_model import (
    SourceFile,
    CodeElement,
    GoFunction,
    GoStructLiteral,
    GoMapLiteral,
    LiteralExpression,
    ReturnExpression,
)

@dataclass(frozen=True)
class Looper:
    index: int
    first: bool
    last: bool


def looper(collection):
    for idx, element in enumerate(collection):
        meta = Looper(
            first=(idx == 0),
            last=(idx == len(collection) - 1),
            index=idx,
        )
        yield meta, element


class Line:

    def __init__(self, writer, start):
        self._writer = writer
        self._start = start

    def __lshift__(self, other):
        self._writer._write(self._start)
        self._writer._write(other)
        return self._writer


class CodeWriter:

    def __init__(self, stream):
        self.stream = stream
        self.indent = 0
    
    def __lshift__(self, other):
        self._write(other)
        return self

    @property
    def line(self):
        indent = "\t" * self.indent
        return Line(self, indent)
    
    def _indent(s):
        s.indent += 1
    
    def _dedent(s):
        s.indent -= 1
    
    def _write(s, stuff):
        s.stream.write(stuff)
    
    def _write_line(s, stuff=""):
        s.line << stuff << "\n"
    
    def write_file(s, source_file: SourceFile):
        s << (source_file.header or "")
        
        if source_file.header:
            s << "\n\n"
        
        s._write_line(f"package {source_file.package_name}")
        
        if source_file.imports:
            s._write_line()
            s._write_line("import (")
            s._indent()
            
            sorted_imports = sorted(source_file.imports, key=lambda x: (x.startswith('github.com'), x))
            
            for i in sorted_imports:
                s._write_line(f'"{i}"')
            
            s._dedent()
            s._write_line(")")
        
        if source_file.declarations:
            s._write_line()
                
            for declaration in source_file.declarations:
                declaration.write(s)
                s << "\n"
        else:
            s << "\n"
    
    
    def write_function(s, function: GoFunction):
        s.line << "func " << function.name << "("
        
        # paramters
        for I, param in looper(function.parameters):
            s << param.name
            
            if param.varargs:
                s << "..."
            else:
                s << " "
            
            
            if param.type:
                s << param.type
                
            if not I.last:
                s << ", "
        
        s << ") "
        
        # return types
        if len(function.return_types) > 1:
            s << "("
        
        for I, return_type in looper(function.return_types):
            s << return_type
            
            if not I.last:
                s << ", "
        
        if len(function.return_types) > 1:
            s << ")"
        
        # function body
        s << " {" << "\n"
        s._indent()
        
        for part in function.body:
            if isinstance(part, str):
                for line in part.splitlines():
                    s.line << line << "\n"
            else:
                part.write(s)
                s << "\n"
        
        s._dedent()
        s.line << "}" << "\n"
        
    
    def write_literal(s, literal: LiteralExpression):
        s << literal.expression
    
        
    def write_struct_literal(s, struct: GoStructLiteral):
        s << struct.type << "{" << "\n"
        s._indent()
        
        
        if struct.aligned is True:
            max_key_length = max([ len(_) for _ in struct.fields.keys() ])
        
        for I, (name,value) in looper(struct.fields.items()):
            
            # a little padding for large objects
            if not I.first:
                if isinstance(value, GoMapLiteral) or isinstance(value, GoStructLiteral):
                    s.line << '\n'

            s.line << name << ":"
            
            if struct.aligned is True:
                space = max_key_length - len(name) + 1
                s << (' ' * space)
            else:
                s << ' '
            
            if isinstance(value, CodeElement):
                value.write(s)
            else:
                s << str(value)
            
            s << ",\n"
        
        s._dedent()
        s.line << "}"
        
    
    def write_map_literal(s, map: GoMapLiteral):
        s << f"map[{map.key_type}]{map.value_type}" << "{\n"
        s._indent()
        
        for I, (name,value) in looper(map.fields.items()):
            s.line << f'"{name}": '
            value.write(s)
            s << ",\n"
        
        s._dedent()
        s.line << "}"
    
    
    def write_return_expression(s, expr: ReturnExpression):
        s.line << "return "
        expr.expression.write(s)


def write_file_to_string(source: SourceFile):
    stream = StringIO()
    writer = CodeWriter(stream)
    writer.write_file(source)
    
    return stream.getvalue()