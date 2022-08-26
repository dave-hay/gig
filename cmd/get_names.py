import os


def main():
    names = []
    path = os.getcwd()

    for _, file in enumerate(os.listdir(path + "/templates")):
        f = file.strip(".txt")
        names.append(f)

    with open("names", "w") as fp:
        fp.write("{")
        for name in names:
            fp.write(f'"{name}",')
        fp.write("}")


if __name__ == "__main__":
    main()
