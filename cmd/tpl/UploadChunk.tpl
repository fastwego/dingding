
func UploadChunk(ctx *dingding.App, media string, params url.Values) (resp []byte, err error) {

	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(media))
		if err != nil {
			return
		}
		file, err := os.Open(media)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

		// field
		err = m.WriteField("agent_id", params.Get("agent_id"))
		if err != nil {
			return
		}

		err = m.WriteField("upload_id", params.Get("upload_id"))
        if err != nil {
            return
        }

        err = m.WriteField("chunk_sequence", params.Get("chunk_sequence"))
        if err != nil {
            return
        }

	}()

    return ctx.Client.HTTPPost(apiUploadChunk, r,  m.FormDataContentType())
    
}