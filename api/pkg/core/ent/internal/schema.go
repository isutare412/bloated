// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/isutare412/bloated/api/pkg/core/ent/schema","Package":"github.com/isutare412/bloated/api/pkg/core/ent","Schemas":[{"name":"BannedIP","config":{"Table":""},"fields":[{"name":"create_time","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"update_time","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"ip","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":50,"unique":true,"validators":2,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"country","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":60,"optional":true,"validators":2,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}}]},{"name":"Todo","config":{"Table":""},"edges":[{"name":"owner","type":"User","field":"owner_id","ref_name":"todos","unique":true,"inverse":true,"required":true}],"fields":[{"name":"create_time","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"update_time","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":36,"optional":true,"validators":2,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"owner_id","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","PkgName":"uuid","Nillable":false,"RType":{"Name":"UUID","Ident":"uuid.UUID","Kind":17,"PkgPath":"github.com/google/uuid","Methods":{"ClockSequence":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Domain":{"In":[],"Out":[{"Name":"Domain","Ident":"uuid.Domain","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"ID":{"In":[],"Out":[{"Name":"uint32","Ident":"uint32","Kind":10,"PkgPath":"","Methods":null}]},"MarshalBinary":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"MarshalText":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"NodeID":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Time":{"In":[],"Out":[{"Name":"Time","Ident":"uuid.Time","Kind":6,"PkgPath":"github.com/google/uuid","Methods":null}]},"URN":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalBinary":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalText":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Variant":{"In":[],"Out":[{"Name":"Variant","Ident":"uuid.Variant","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"Version":{"In":[],"Out":[{"Name":"Version","Ident":"uuid.Version","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]}}}},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"task","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":3000,"validators":2,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}]},{"name":"User","config":{"Table":""},"edges":[{"name":"todos","type":"Todo","annotations":{"EntSQL":{"on_delete":"CASCADE"}}}],"fields":[{"name":"id","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","PkgName":"uuid","Nillable":false,"RType":{"Name":"UUID","Ident":"uuid.UUID","Kind":17,"PkgPath":"github.com/google/uuid","Methods":{"ClockSequence":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Domain":{"In":[],"Out":[{"Name":"Domain","Ident":"uuid.Domain","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"ID":{"In":[],"Out":[{"Name":"uint32","Ident":"uint32","Kind":10,"PkgPath":"","Methods":null}]},"MarshalBinary":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"MarshalText":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"NodeID":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Time":{"In":[],"Out":[{"Name":"Time","Ident":"uuid.Time","Kind":6,"PkgPath":"github.com/google/uuid","Methods":null}]},"URN":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalBinary":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalText":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Variant":{"In":[],"Out":[{"Name":"Variant","Ident":"uuid.Variant","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"Version":{"In":[],"Out":[{"Name":"Version","Ident":"uuid.Version","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"email","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":256,"validators":2,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"user_name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":800,"validators":2,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"given_name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":800,"validators":2,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"family_name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":800,"validators":2,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"photo_url","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"origin","type":{"Type":6,"Ident":"enum.Issuer","PkgPath":"github.com/isutare412/bloated/api/pkg/core/enum","PkgName":"enum","Nillable":false,"RType":{"Name":"Issuer","Ident":"enum.Issuer","Kind":24,"PkgPath":"github.com/isutare412/bloated/api/pkg/core/enum","Methods":{"Values":{"In":[],"Out":[{"Name":"","Ident":"[]string","Kind":23,"PkgPath":"","Methods":null}]}}}},"enums":[{"N":"none","V":"none"},{"N":"google","V":"google"}],"position":{"Index":6,"MixedIn":false,"MixinIndex":0}}]}],"Features":["schema/snapshot"]}`
