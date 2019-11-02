from dataclasses import dataclass

@dataclass
class ResourceAttribute:
    name: str
    type: str
    element: str
    required: bool
    will_replace: bool
