package SkyLine_Backend

func init() {
	RegisterBuiltin("crypt.hash", func(env *Environment_of_environment, args ...Object) Object {
		return (Crypto_Hasher(args...))
	})
}
