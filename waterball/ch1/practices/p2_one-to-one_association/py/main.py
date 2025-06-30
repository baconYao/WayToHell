from fruit.fruit import Fruit
from hero.hero import Hero
from pet.pet import Pet


def main():
    # Create a new Hero instance
    h = Hero()

    # Create a new LevelSheet instance
    p = Pet(name="Cat")
    h.pet = p

    print(f"Hero 目前血量: {h.hp}")
    print(f"Hero's 寵物名稱: {h.pet.name}")

    for i in range(5):
        p.eat_fruit(Fruit())

    h.remove_pet()

    for i in range(5):
        p.eat_fruit(Fruit())


# for i := 0; i < 5; i++ {
# 	pet.EatFruit(fruit.NewFruit())
# }

# h.RemovePet()

# for i := 0; i < 5; i++ {
# 	pet.EatFruit(fruit.NewFruit())
# }


if __name__ == "__main__":
    main()
