package maul

func ExportNewRepository(owner, name string) repository {
	return repository{
		owner: owner,
		name:  name,
	}
}
