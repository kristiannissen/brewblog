module brewblog

go 1.21.4

replace brewblog/api/entry => ./api/entry

require brewblog/api/entry v0.1.0

require github.com/rpdg/vercel_blob v0.1.0 // indirect
