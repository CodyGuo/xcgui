package doc

/*
Win32 C/C++ golang 字符对照表

	WIN32类型		C/C++ 类型			GO 类型
	HANDLE			void *				uintptr
	BYTE			unsigned char		uint8, byte
	SHORT			short				int16
	WORD			unsigned short		uint16
	INT				int					int32, int
	UINT			unsigned int		uint32
	LONG			long				int32
	BOOL			int					int
	DWORD			unsigned long		uint32
	ULONG			unsigned long		uint32
	CHAR			char				byte
	WCHAR			wchar_t				uint16
	LPSTR			utf8/char *			*byte
	LPCSTR			const utf8/char *	*byte, syscall.StringBytePtr(), xc.UTF8PtrToSting()
	LPWSTR			wchar_t *			*uint16
	LPCWSTR			const wchar_t *		*uint16, syscall.StringToUTF16Ptr()
	FLOAT			float				float32
	DOUBLE			double				float64
	LONGLONG		__int64				int64
	DWORD64			unsigned __int64	uint64
*/
