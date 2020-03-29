import functools

def doc(f):
	@functools.wraps(f)
	def inner(*args, **kwargs):
		print(f.__doc__)
		return f(*args, **kwargs)
	return inner
	
def log(log_file):
	def inner(f):
		@functools.wraps(f)
		def wrap(*args, **kwargs):
			r = f(*args, **kwargs)
			with open(log_file, "a") as fp:
				fp.write(str(args))
				fp.write(str(kwargs))
				fp.write(str(r))
				fp.write("\n")
			return r
		return wrap
	return inner

@doc
@log("add.log")
def add(a,b,*args, **kwargs):
	'''add a and b'''
	return a+b

@doc
@log("redu.log")
def redu(a,b,*args, **kwargs):
	'''redu a and b'''
	return a-b


add(1,2)

redu(2,1)