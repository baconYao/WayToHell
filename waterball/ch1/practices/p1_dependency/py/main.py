from hero.hero import Hero
from hero.levelsheet import LevelSheet


def main():
    # Create a new Hero instance
    h = Hero()

    # Create a new LevelSheet instance
    ls = LevelSheet()

    # Test gaining experience
    exps = [0, 100, 900, -200]
    for exp in exps:
        try:
            h.gain_exp(exp, ls)
        except ValueError as e:
            print(e)


if __name__ == "__main__":
    main()
