package module6_2

// PostCount - count posts
func PostCount(client Client) (int, error) {
	posts, err := client.GetList()
	if err != nil {
		return 0, err
	}

	return len(posts), nil
}
