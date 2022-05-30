#!/usr/bin/env python3
import sys, os

package = sys.stdin.readline().strip()

res: list[str] = []

for field in sys.stdin.readlines():
    v = field.strip()
    if not v:
        continue

    f, rest = v.split(" ", 1)
    header = f"{f} (Default){rest} {{"

    fname, r = rest.split("(", 1)
    argsWithTypes, rest = r.split(")", 1)
    argsWithTypes = argsWithTypes.split(", ")
    args = [arg.split()[0] for arg in argsWithTypes if arg]

    joinedArgs = ", ".join(args)
    ret = "return " if rest.strip() else ""
    body = f"{ret}{package}.{fname}({joinedArgs})"

    res.append(f"{header}\n\t{body}\n}}")

print(f"package {package}")
print(f'import "{package}"')
print("\n".join(res))
