from enum import Enum, auto, unique
from collections import defaultdict
from dataclasses import dataclass, asdict, field

from unquietcode.tools.cfn_provider.utils import snake_caps


class DataDict(dict):
    
    def as_dict(self):
        return self


@dataclass
class Package:
    name: str = ""
    parent: any = None
    resources: dict = field(default_factory=lambda: defaultdict(lambda: DataDict()))
    datasources: dict = field(default_factory=lambda: defaultdict(lambda: DataDict()))
    properties: dict = field(default_factory=lambda: defaultdict(lambda: DataDict()))
    subpackages: dict = field(default_factory=lambda: {})

    @property
    def full_path(self):
        path = ""
        p = self
        
        while (p is not None):
            path = p.name + '/' + path
            p = p.parent
        
        return path[0:-1]
        
    def as_dict(self):
        return {
            'name': self.name,
            'properties': {k:v.as_dict() for k,v in self.properties.items()},
            'resources': {k:v.as_dict() for k,v in self.resources.items()},
            'subpackages': {k:v.as_dict() for k,v in self.subpackages.items()},
        }


@unique
class CF_Type(Enum):
    def _generate_next_value_(name, start, count, last_values):
        return name
    
    # primitives
    String     = auto()
    Long       = auto()
    Integer    = auto()
    Double     = auto()
    Boolean    = auto()
    Timestamp  = auto()
    Json       = auto()
    
    # collections
    List       = auto()
    Map        = auto()


@unique
class TF_Type(Enum):
    def _generate_next_value_(name, start, count, last_values):
        return f'schema.Type{name}'
            
    String  = auto()
    Int     = auto()
    Float   = auto()
    Bool    = auto()
    Map     = auto()
    List    = auto()
    Set     = auto()
    
    @property
    def schema_type(self):
        return "&schema.Schema{Type: "+self.value+"}"

    def __repr__(self):
        return self.value
        
    def __str__(self):
        return self.value

@dataclass
class AttributeType:
    type: str
    element_type: str
    min_items: int
    max_items: int
    set_function: str
    
    def as_dict(self):
        return asdict(self)

    
@dataclass
class ResourceAttribute:
    name: str
    type: AttributeType
    required: bool
    will_replace: bool

    @property
    def go_symbol(self):
        return snake_caps(self.name)
        
    def as_dict(self):
        return asdict(self)


@dataclass
class ComplexType:
    name: str = ""
    package: Package = None

    @property
    def go_symbol(self):
        return snake_caps(self.name)


@dataclass
class Resource(ComplexType):
    cfn_type: str = ""
    attributes: dict = field(default_factory=lambda: {})

    def as_dict(self):
        return asdict(self)


@dataclass
class Property(ComplexType):
    attributes: dict = field(default_factory=lambda: {})
    
    # TODO deprecated?
    resource_name: str = ""
    
    def as_dict(self):
        return asdict(self)


@dataclass
class Provider:
    top_level_package: Package
    cfn_version: str