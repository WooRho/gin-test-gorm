package metadata

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func SetMDToIncoming(ctx context.Context, k, v string) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)
	newMd := metadata.MD{}
	for key, value := range md {
		newMd[key] = value
	}
	newMd.Set(k, v)
	return metadata.NewIncomingContext(ctx, newMd)
}

func GetMD(ctx context.Context, k string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	v := md.Get(k)
	if len(v) <= 0 {
		return ""
	}

	return v[0]
}
