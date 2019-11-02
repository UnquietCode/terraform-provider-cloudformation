from collections import defaultdict
from dataclasses import dataclass, asdict, field


class DataDict(dict):
    
    def as_dict(self):
        return self


@dataclass
class Package:
    name: str = ""
    parent: any = None
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
    
    
@dataclass
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
    attributes: dict = field(default_factory=lambda: {})
    
    def as_dict(self):
        return asdict(self)


@dataclass(frozen=True)
class Property:
    name: str
    attributes: dict
    
    package_name: str
    resource_name: str

    def as_dict(self):
        return asdict(self)
