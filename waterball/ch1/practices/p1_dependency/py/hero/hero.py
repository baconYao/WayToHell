from .levelsheet import LevelSheet


class Hero:
    def __init__(self):
        self._level = 1
        self._total_exp = 0
        self._hp = 100

    def gain_exp(self, exp: int, level_sheet: LevelSheet):
        if exp < 0:
            raise ValueError(f"error: 無法獲得負數的經驗值 {exp}")
        current_level = self.level
        self.total_exp = self.total_exp + exp
        self.level = level_sheet.query_level(self.total_exp)
        print(
            f"英雄目前等級 {current_level}，獲得 {exp} EXP，最新總共經驗值為 {self.total_exp}，最新等級為 {self.level}"
        )

    @property
    def hp(self) -> int:
        return self._hp

    @hp.setter
    def hp(self, new_hp: int):
        if new_hp < 0:
            raise ValueError(f"HP must be greater than or equal to 0, got {new_hp}")
        self._hp = new_hp

    @property
    def total_exp(self) -> int:
        return self._total_exp

    @total_exp.setter
    def total_exp(self, exp: int):
        if exp < 0:
            raise ValueError(f"Total EXP must be greater than or equal to 0, got {exp}")
        self._total_exp = exp

    @property
    def level(self) -> int:
        return self._level

    @level.setter
    def level(self, new_level: int):
        if new_level < 0:
            raise ValueError(
                f"Level must be greater than or equal to 0, got {new_level}"
            )
        self._level = new_level
