#!/usr/bin/env python3
import sys


def clean(line: str) -> str:
    return remove_comment(line).strip()


def remove_comment(line: str) -> str:
    return line.split("//")[0]


package = sys.stdin.readline().strip()

res: list[str] = []

for var in sys.stdin.readlines():
    v = clean(var)
    if v:
        name, t = v.split()
        res.append(f"func (Default){name}() {t}{{\n\treturn {package}.{name}\n}}")

print(f"package {package}")
print(f'import "{package}"')
print("\n".join(res))
