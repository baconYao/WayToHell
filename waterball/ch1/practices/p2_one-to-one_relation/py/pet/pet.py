from abc import abstractmethod

from fruit.fruit import Fruit


class Owner:
    @property
    @abstractmethod
    def hp(self):
        pass

    @hp.setter
    @abstractmethod
    def hp(self, hp):
        pass


class Pet:
    def __init__(self, name: str):
        self._name = name
        self._owner = None

    @property
    def name(self) -> str:
        return self._name

    @name.setter
    def name(self, name: str):
        if name == "":
            raise ValueError("Error: cannot set the name as empty string")
        self._name = name

    @property
    def owner(self) -> Owner:
        return self._owner

    @owner.setter
    def owner(self, owner: Owner):
        self._owner = owner

    def eat_fruit(self, fruit: Fruit):
        print("寵物吃水果...")

        if self._owner:
            self._owner.hp = self._owner.hp + 10
