from io import StringIO
from unquietcode.tools.cfn_provider.golang.code_model import (
    SourceFile,
    GoFunction,
)

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
            
            for i in source_file.imports:
                s._write_line(f'"{i}"')
            
            s._dedent()
            s._write_line(")")
        
        if source_file.declarations:
            s._write_line()
                
            for declaration in source_file.declarations:
                declaration.write(s)
    
    
    def write_function(s, function: GoFunction):
        s.line << "func " << function.name << "("
        
        # paramters
        for idx, param in enumerate(function.parameters):
            s << param.name
            
            if param.varargs:
                s << "..."
            else:
                s << " "
            
            
            if param.type:
                s << param.type
                
            if idx < len(function.parameters) - 1:
                s << ", "
        
        s << ") "
        
        # return types
        if len(function.return_types) > 1:
            s << "("
        
        for idx, return_type in enumerate(function.return_types):
            s << return_type
            
            if idx < len(function.return_types) - 1:
                s << ", "
        
        if len(function.return_types) > 1:
            s << ")"
        
        # function body
        s << " {" << "\n"
        s._indent()
        
        for line in function.body.splitlines():
            s.line << line << "\n"
        
        s._dedent()
        s.line << "}" << "\n"
        


def write_file_to_string(source: SourceFile):
    stream = StringIO()
    writer = CodeWriter(stream)
    writer.write_file(source)
    
    return stream.getvalue()