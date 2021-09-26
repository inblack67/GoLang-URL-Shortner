package shortner

import "time"

var dummies TRedirects = TRedirects{
	&MRedirect{
		Code:      "ok",
		URL:       "some",
		CreatedAt: time.Now().String(),
	},
	&MRedirect{
		Code:      "ok",
		URL:       "some",
		CreatedAt: time.Now().String(),
	},
	&MRedirect{
		Code:      "ok",
		URL:       "some",
		CreatedAt: time.Now().String(),
	},
	&MRedirect{
		Code:      "ok",
		URL:       "some",
		CreatedAt: time.Now().String(),
	},
	&MRedirect{
		Code:      "ok",
		URL:       "some",
		CreatedAt: time.Now().String(),
	},
}
