# Accounts package
I am using this site as a reference for working with clientv3 for etcd api:
`https://programmer.help/blogs/a-concise-tutorial-of-golang-etcd.html`

```shell
etcdctl get / --prefix
/test/key1
Hello etcd!
/test/key2
value2

etcdctl get / prefix
/test/key1
Hello etcd!
/test/key2
value2
key1
value1
key2
value2
```
