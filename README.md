# nodin

## syntax (uptil now)

```
pkg node;

pub type struct: Node {
	data: pub int;
	next: *node;

	Next: pub func() *node;
	Manipulator: pub func(a: int);
};

(n node)::Next() {
	return n.next;
}

(n *node)::Manipulator(a: int) {
	n.data ^= a;
}
```
