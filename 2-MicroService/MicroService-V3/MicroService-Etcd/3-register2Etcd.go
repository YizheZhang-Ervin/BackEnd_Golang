func NewService() *Service {
	config := clientv3.Config{
		Endpoints:   []string{"106.12.72.181:23791", "106.12.72.181:23792"},
		DialTimeout: 10 * time.Second,
	}
	client, _ := clientv3.New(config)
	return &Service{client: client}
}

func (this *Service) RegService(id string, name string, address string) error {
	fmt.Println(22)
	kv := clientv3.NewKV(this.client)
	key_prefix := "/services/"
	_, err := kv.Put(context.Background(), key_prefix+id+"/"+name, address)
	fmt.Println(err)
	return err
}