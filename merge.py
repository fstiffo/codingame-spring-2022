import re

merged_code = ""
imported_modules = set([])
files_to_import = ["types", "utils", "topology", "bot", "debug",
                   "main", ]
for filename in files_to_import:
    print(filename)
    with open(f"{filename}.go", "r") as f:
        data = f.read()
        data = data.replace("package main", "")

        # single import
        m = re.search(r"import \"([^\"]+)\"", data)
        if m is not None:
            imported_modules.add(m.groups()[0])
            data = re.sub(r"import \"([^\"]+)\"", "", data)
        else:
            # multiple import
            m = re.search(r"import \(\n([^)]+)\n\)", data)
            if m is not None:
                modules = list(map(str.lstrip, m.groups()[0].split("\n")))
                modules = [s.replace('"', '') for s in modules]
                imported_modules.update(modules)
                data = re.sub(r"import \(\n[^)]+\n\)", "", data)

        merged_code += f"\n// ---------------------------------------------\n// {filename}.go\n// ---------------------------------------------" + data

imported_modules_str = "import (\n"
for m in imported_modules:
    imported_modules_str += f'"{m}"' + "\n"
imported_modules_str += ")\n"

merged_code = "package main\n\n" + imported_modules_str + merged_code

with open("codingame-spring-2022.merged.go", "w") as f:
    f.write(merged_code)
