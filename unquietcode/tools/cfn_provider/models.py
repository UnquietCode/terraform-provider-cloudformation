from collections import defaultdict
from dataclasses import dataclass, asdict, field


@dataclass(frozen=True)
class ResourceAttribute:
    name: str
    type: str
    element: str
    required: bool
    will_replace: bool

    def as_dict(self):
        return asdict(self)

class DataDict(dict):
    
    def as_dict(self):
        return self


@dataclass
class Package:
    name: str = ""
    resources: dict = field(default_factory=lambda: defaultdict(lambda: DataDict()))
    properties: dict = field(default_factory=lambda: defaultdict(lambda: DataDict()))
    subpackages: dict = field(default_factory=lambda: {})
    
    def as_dict(self):
        return {
            'name': self.name,
            'properties': {k:v.as_dict() for k,v in self.properties.items()},
            'resources': {k:v.as_dict() for k,v in self.resources.items()},
            'subpackages': {k:v.as_dict() for k,v in self.subpackages.items()},
        }
    
    
@dataclass(frozen=True)
class ResourceAttribute:
    name: str
    type: str
    element: str
    required: bool
    will_replace: bool

    def as_dict(self):
        return asdict(self)

@dataclass
class Resource:
    name: str = ""
    file_name: str = ""
    attributes: dict = field(default_factory=lambda: {})
    
    def as_dict(self):
        return asdict(self)

@dataclass
class Property:
    name: str
    type: str
    package_name: str
    resource_name: str

    def as_dict(self):
        return asdict(self)
    
