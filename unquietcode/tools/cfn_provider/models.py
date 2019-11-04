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
    
    
@dataclass
class ResourceAttribute:
    name: str
    type: str
    element: str
    required: bool
    will_replace: bool
    repeatable: bool

    @property
    def go_symbol(self):
        return snake_caps(self.name)
        
    def as_dict(self):
        return asdict(self)


@dataclass
class Resource:
    package: Package = None
    name: str = ""
    cfn_type: str = ""
    attributes: dict = field(default_factory=lambda: {})

    @property
    def go_symbol(self):
        return snake_caps(self.name)
    
    def as_dict(self):
        return asdict(self)


@dataclass(frozen=True)
class Property:
    name: str
    attributes: dict
    package: Package
    resource_name: str

    @property
    def go_symbol(self):
        return snake_caps(self.name)
    
    def as_dict(self):
        return asdict(self)
