package acl

type aclAuth struct {
	config *ACLConfig
}

func Init() *aclAuth {
	return New("./plugins/auth/authfile/acl.conf")
}

func New(file string)*aclAuth  {
	aclConfig, err := AclConfigLoad(file)
	if err != nil {
		panic(err)
	}
	return &aclAuth{
		config: aclConfig,
	}
}

func (a *aclAuth) CheckConnect(clientID, username, password string) bool {
	return true
}

func (a *aclAuth) CheckACL(action, clientID, username, ip, topic string) bool {
	return checkTopicAuth(a.config, action, username, ip, clientID, topic)
}
