OUTDIR = "/home/jeetu/Desktop"	

def defineAst( baseName : str, types_n ):
	path = OUTDIR + "/" + baseName + ".go"
	with open(path, 'w') as fp:
		fp.write("package gox\n")
		fp.write("\n")


		for s_type in types_n:
			className = s_type.split(':')[0].strip()
			fields = s_type.split(':')[1].strip()
			
			defineType(fp, baseName, className, fields)

def defineType(fp, baseName, className, fields):
	_fields = fields.split(",")

	fp.write(f"type {className} struct {{\n")
	fp.write()




defineAst("Expr", [
	"Binary   : Expr left, Token operator, Expr right",
    "Grouping : Expr expression",
    "Literal  : Object value",
    "Unary    : Token operator, Expr right"
	])

			

