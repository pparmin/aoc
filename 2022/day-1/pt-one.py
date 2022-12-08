def read_file(input: str) -> list[str]:
    lines = []
    with open(input) as f: 
        lines = f.readlines()
    return lines

def calc_elves(elves: list[str]): 
    for e in elves:
        print(e)


if __name__ == '__main__':
    input = "input.txt"
    read_file(input)
