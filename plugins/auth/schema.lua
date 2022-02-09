return {
    name = "auth",
    fields = {
        {
            config = {
                type = "record",
                fields = {
                    {
                        request_header = {
                            type = "string",
                            required = false,
                            default = "Authorization"
                        },
                    },
                    {
                        forward_header = {
                            type = "string",
                            required = false,
                            default = "X-Authenticated-UID"
                        },
                    },
                    {
                        default_uid = {
                            type = "string",
                            required = true
                        }
                    },
                }
            }
        }
    }
}