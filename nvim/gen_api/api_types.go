package main

type API struct {
	NvimBufAddDecoration    APIBufAddDecoration    `json:"nvim__buf_add_decoration" msgpack:"nvim_buf_add_decoration"`
	NvimBufRedrawRange      APIBufRedrawRange      `json:"nvim__buf_redraw_range" msgpack:"nvim_buf_redraw_range"`
	NvimBufSetLuahl         APIBufSetLuahl         `json:"nvim__buf_set_luahl" msgpack:"nvim_buf_set_luahl"`
	NvimBufStats            APIBufStats            `json:"nvim__buf_stats" msgpack:"nvim_buf_stats"`
	NvimGetLibDir           APIGetLibDir           `json:"nvim__get_lib_dir" msgpack:"nvim_get_lib_dir"`
	NvimID                  APIID                  `json:"nvim__id" msgpack:"nvim_id"`
	NvimIDArray             APIIDArray             `json:"nvim__id_array" msgpack:"nvim_id_array"`
	NvimIDDictionary        APIIDDictionary        `json:"nvim__id_dictionary" msgpack:"nvim_id_dictionary"`
	NvimIDFloat             APIIDFloat             `json:"nvim__id_float" msgpack:"nvim_id_float"`
	NvimInspectCell         APIInspectCell         `json:"nvim__inspect_cell" msgpack:"nvim_inspect_cell"`
	NvimPutAttr             APIPutAttr             `json:"nvim__put_attr" msgpack:"nvim_put_attr"`
	NvimStats               APIStats               `json:"nvim__stats" msgpack:"nvim_stats"`
	NvimBufAddHighlight     APIBufAddHighlight     `json:"nvim_buf_add_highlight" msgpack:"nvim_buf_add_highlight"`
	NvimBufAttach           APIBufAttach           `json:"nvim_buf_attach" msgpack:"nvim_buf_attach"`
	NvimBufClearNamespace   APIBufClearNamespace   `json:"nvim_buf_clear_namespace" msgpack:"nvim_buf_clear_namespace"`
	NvimBufDelExtmark       APIBufDelExtmark       `json:"nvim_buf_del_extmark" msgpack:"nvim_buf_del_extmark"`
	NvimBufDelKeymap        APIBufDelKeymap        `json:"nvim_buf_del_keymap" msgpack:"nvim_buf_del_keymap"`
	NvimBufDelVar           APIBufDelVar           `json:"nvim_buf_del_var" msgpack:"nvim_buf_del_var"`
	NvimBufDetach           APIBufDetach           `json:"nvim_buf_detach" msgpack:"nvim_buf_detach"`
	NvimBufGetChangedtick   APIBufGetChangedtick   `json:"nvim_buf_get_changedtick" msgpack:"nvim_buf_get_changedtick"`
	NvimBufGetCommands      APIBufGetCommands      `json:"nvim_buf_get_commands" msgpack:"nvim_buf_get_commands"`
	NvimBufGetExtmarkByID   APIBufGetExtmarkByID   `json:"nvim_buf_get_extmark_by_id" msgpack:"nvim_buf_get_extmark_by_id"`
	NvimBufGetExtmarks      APIBufGetExtmarks      `json:"nvim_buf_get_extmarks" msgpack:"nvim_buf_get_extmarks"`
	NvimBufGetKeymap        APIBufGetKeymap        `json:"nvim_buf_get_keymap" msgpack:"nvim_buf_get_keymap"`
	NvimBufGetLines         APIBufGetLines         `json:"nvim_buf_get_lines" msgpack:"nvim_buf_get_lines"`
	NvimBufGetMark          APIBufGetMark          `json:"nvim_buf_get_mark" msgpack:"nvim_buf_get_mark"`
	NvimBufGetName          APIBufGetName          `json:"nvim_buf_get_name" msgpack:"nvim_buf_get_name"`
	NvimBufGetOffset        APIBufGetOffset        `json:"nvim_buf_get_offset" msgpack:"nvim_buf_get_offset"`
	NvimBufGetOption        APIBufGetOption        `json:"nvim_buf_get_option" msgpack:"nvim_buf_get_option"`
	NvimBufGetVar           APIBufGetVar           `json:"nvim_buf_get_var" msgpack:"nvim_buf_get_var"`
	NvimBufGetVirtualText   APIBufGetVirtualText   `json:"nvim_buf_get_virtual_text" msgpack:"nvim_buf_get_virtual_text"`
	NvimBufIsLoaded         APIBufIsLoaded         `json:"nvim_buf_is_loaded" msgpack:"nvim_buf_is_loaded"`
	NvimBufIsValid          APIBufIsValid          `json:"nvim_buf_is_valid" msgpack:"nvim_buf_is_valid"`
	NvimBufLineCount        APIBufLineCount        `json:"nvim_buf_line_count" msgpack:"nvim_buf_line_count"`
	NvimBufSetExtmark       APIBufSetExtmark       `json:"nvim_buf_set_extmark" msgpack:"nvim_buf_set_extmark"`
	NvimBufSetKeymap        APIBufSetKeymap        `json:"nvim_buf_set_keymap" msgpack:"nvim_buf_set_keymap"`
	NvimBufSetLines         APIBufSetLines         `json:"nvim_buf_set_lines" msgpack:"nvim_buf_set_lines"`
	NvimBufSetName          APIBufSetName          `json:"nvim_buf_set_name" msgpack:"nvim_buf_set_name"`
	NvimBufSetOption        APIBufSetOption        `json:"nvim_buf_set_option" msgpack:"nvim_buf_set_option"`
	NvimBufSetVar           APIBufSetVar           `json:"nvim_buf_set_var" msgpack:"nvim_buf_set_var"`
	NvimBufSetVirtualText   APIBufSetVirtualText   `json:"nvim_buf_set_virtual_text" msgpack:"nvim_buf_set_virtual_text"`
	NvimCallAtomic          APICallAtomic          `json:"nvim_call_atomic" msgpack:"nvim_call_atomic"`
	NvimCallDictFunction    APICallDictFunction    `json:"nvim_call_dict_function" msgpack:"nvim_call_dict_function"`
	NvimCallFunction        APICallFunction        `json:"nvim_call_function" msgpack:"nvim_call_function"`
	NvimCommand             APICommand             `json:"nvim_command" msgpack:"nvim_command"`
	NvimCreateBuf           APICreateBuf           `json:"nvim_create_buf" msgpack:"nvim_create_buf"`
	NvimCreateNamespace     APICreateNamespace     `json:"nvim_create_namespace" msgpack:"nvim_create_namespace"`
	NvimDelCurrentLine      APIDelCurrentLine      `json:"nvim_del_current_line" msgpack:"nvim_del_current_line"`
	NvimDelKeymap           APIDelKeymap           `json:"nvim_del_keymap" msgpack:"nvim_del_keymap"`
	NvimDelVar              APIDelVar              `json:"nvim_del_var" msgpack:"nvim_del_var"`
	NvimErrWrite            APIErrWrite            `json:"nvim_err_write" msgpack:"nvim_err_write"`
	NvimErrWriteln          APIErrWriteln          `json:"nvim_err_writeln" msgpack:"nvim_err_writeln"`
	NvimEval                APIEval                `json:"nvim_eval" msgpack:"nvim_eval"`
	NvimExec                APIExec                `json:"nvim_exec" msgpack:"nvim_exec"`
	NvimExecLua             APIExecLua             `json:"nvim_exec_lua" msgpack:"nvim_exec_lua"`
	NvimFeedkeys            APIFeedkeys            `json:"nvim_feedkeys" msgpack:"nvim_feedkeys"`
	NvimGetAPIInfo          APIGetAPIInfo          `json:"nvim_get_api_info" msgpack:"nvim_get_api_info"`
	NvimGetChanInfo         APIGetChanInfo         `json:"nvim_get_chan_info" msgpack:"nvim_get_chan_info"`
	NvimGetColorByName      APIGetColorByName      `json:"nvim_get_color_by_name" msgpack:"nvim_get_color_by_name"`
	NvimGetColorMap         APIGetColorMap         `json:"nvim_get_color_map" msgpack:"nvim_get_color_map"`
	NvimGetCommands         APIGetCommands         `json:"nvim_get_commands" msgpack:"nvim_get_commands"`
	NvimGetContext          APIGetContext          `json:"nvim_get_context" msgpack:"nvim_get_context"`
	NvimGetCurrentBuf       APIGetCurrentBuf       `json:"nvim_get_current_buf" msgpack:"nvim_get_current_buf"`
	NvimGetCurrentLine      APIGetCurrentLine      `json:"nvim_get_current_line" msgpack:"nvim_get_current_line"`
	NvimGetCurrentTabpage   APIGetCurrentTabpage   `json:"nvim_get_current_tabpage" msgpack:"nvim_get_current_tabpage"`
	NvimGetCurrentWin       APIGetCurrentWin       `json:"nvim_get_current_win" msgpack:"nvim_get_current_win"`
	NvimGetHlByID           APIGetHlByID           `json:"nvim_get_hl_by_id" msgpack:"nvim_get_hl_by_id"`
	NvimGetHlByName         APIGetHlByName         `json:"nvim_get_hl_by_name" msgpack:"nvim_get_hl_by_name"`
	NvimGetHlIDByName       APIGetHlIDByName       `json:"nvim_get_hl_id_by_name" msgpack:"nvim_get_hl_id_by_name"`
	NvimGetKeymap           APIGetKeymap           `json:"nvim_get_keymap" msgpack:"nvim_get_keymap"`
	NvimGetMode             APIGetMode             `json:"nvim_get_mode" msgpack:"nvim_get_mode"`
	NvimGetNamespaces       APIGetNamespaces       `json:"nvim_get_namespaces" msgpack:"nvim_get_namespaces"`
	NvimGetOption           APIGetOption           `json:"nvim_get_option" msgpack:"nvim_get_option"`
	NvimGetProc             APIGetProc             `json:"nvim_get_proc" msgpack:"nvim_get_proc"`
	NvimGetProcChildren     APIGetProcChildren     `json:"nvim_get_proc_children" msgpack:"nvim_get_proc_children"`
	NvimGetRuntimeFile      APIGetRuntimeFile      `json:"nvim_get_runtime_file" msgpack:"nvim_get_runtime_file"`
	NvimGetVar              APIGetVar              `json:"nvim_get_var" msgpack:"nvim_get_var"`
	NvimGetVvar             APIGetVvar             `json:"nvim_get_vvar" msgpack:"nvim_get_vvar"`
	NvimInput               APIInput               `json:"nvim_input" msgpack:"nvim_input"`
	NvimInputMouse          APIInputMouse          `json:"nvim_input_mouse" msgpack:"nvim_input_mouse"`
	NvimListBufs            APIListBufs            `json:"nvim_list_bufs" msgpack:"nvim_list_bufs"`
	NvimListChans           APIListChans           `json:"nvim_list_chans" msgpack:"nvim_list_chans"`
	NvimListRuntimePaths    APIListRuntimePaths    `json:"nvim_list_runtime_paths" msgpack:"nvim_list_runtime_paths"`
	NvimListTabpages        APIListTabpages        `json:"nvim_list_tabpages" msgpack:"nvim_list_tabpages"`
	NvimListUis             APIListUis             `json:"nvim_list_uis" msgpack:"nvim_list_uis"`
	NvimListWins            APIListWins            `json:"nvim_list_wins" msgpack:"nvim_list_wins"`
	NvimLoadContext         APILoadContext         `json:"nvim_load_context" msgpack:"nvim_load_context"`
	NvimOpenWin             APIOpenWin             `json:"nvim_open_win" msgpack:"nvim_open_win"`
	NvimOutWrite            APIOutWrite            `json:"nvim_out_write" msgpack:"nvim_out_write"`
	NvimParseExpression     APIParseExpression     `json:"nvim_parse_expression" msgpack:"nvim_parse_expression"`
	NvimPaste               APIPaste               `json:"nvim_paste" msgpack:"nvim_paste"`
	NvimPut                 APIPut                 `json:"nvim_put" msgpack:"nvim_put"`
	NvimReplaceTermcodes    APIReplaceTermcodes    `json:"nvim_replace_termcodes" msgpack:"nvim_replace_termcodes"`
	NvimSelectPopupmenuItem APISelectPopupmenuItem `json:"nvim_select_popupmenu_item" msgpack:"nvim_select_popupmenu_item"`
	NvimSetClientInfo       APISetClientInfo       `json:"nvim_set_client_info" msgpack:"nvim_set_client_info"`
	NvimSetCurrentBuf       APISetCurrentBuf       `json:"nvim_set_current_buf" msgpack:"nvim_set_current_buf"`
	NvimSetCurrentDir       APISetCurrentDir       `json:"nvim_set_current_dir" msgpack:"nvim_set_current_dir"`
	NvimSetCurrentLine      APISetCurrentLine      `json:"nvim_set_current_line" msgpack:"nvim_set_current_line"`
	NvimSetCurrentTabpage   APISetCurrentTabpage   `json:"nvim_set_current_tabpage" msgpack:"nvim_set_current_tabpage"`
	NvimSetCurrentWin       APISetCurrentWin       `json:"nvim_set_current_win" msgpack:"nvim_set_current_win"`
	NvimSetKeymap           APISetKeymap           `json:"nvim_set_keymap" msgpack:"nvim_set_keymap"`
	NvimSetOption           APISetOption           `json:"nvim_set_option" msgpack:"nvim_set_option"`
	NvimSetVar              APISetVar              `json:"nvim_set_var" msgpack:"nvim_set_var"`
	NvimSetVvar             APISetVvar             `json:"nvim_set_vvar" msgpack:"nvim_set_vvar"`
	NvimStrwidth            APIStrwidth            `json:"nvim_strwidth" msgpack:"nvim_strwidth"`
	NvimSubscribe           APISubscribe           `json:"nvim_subscribe" msgpack:"nvim_subscribe"`
	NvimTabpageDelVar       APITabpageDelVar       `json:"nvim_tabpage_del_var" msgpack:"nvim_tabpage_del_var"`
	NvimTabpageGetNumber    APITabpageGetNumber    `json:"nvim_tabpage_get_number" msgpack:"nvim_tabpage_get_number"`
	NvimTabpageGetVar       APITabpageGetVar       `json:"nvim_tabpage_get_var" msgpack:"nvim_tabpage_get_var"`
	NvimTabpageGetWin       APITabpageGetWin       `json:"nvim_tabpage_get_win" msgpack:"nvim_tabpage_get_win"`
	NvimTabpageIsValid      APITabpageIsValid      `json:"nvim_tabpage_is_valid" msgpack:"nvim_tabpage_is_valid"`
	NvimTabpageListWins     APITabpageListWins     `json:"nvim_tabpage_list_wins" msgpack:"nvim_tabpage_list_wins"`
	NvimTabpageSetVar       APITabpageSetVar       `json:"nvim_tabpage_set_var" msgpack:"nvim_tabpage_set_var"`
	NvimUIAttach            APIUIAttach            `json:"nvim_ui_attach" msgpack:"nvim_ui_attach"`
	NvimUIDetach            APIUidetach            `json:"nvim_ui_detach" msgpack:"nvim_ui_detach"`
	NvimUIPumSetHeight      APIUipumSetHeight      `json:"nvim_ui_pum_set_height" msgpack:"nvim_ui_pum_set_height"`
	NvimUISetOption         APIUISetOption         `json:"nvim_ui_set_option" msgpack:"nvim_ui_set_option"`
	NvimUITryResize         APIUITryResize         `json:"nvim_ui_try_resize" msgpack:"nvim_ui_try_resize"`
	NvimUITryResizeGrid     APIUITryResizeGrid     `json:"nvim_ui_try_resize_grid" msgpack:"nvim_ui_try_resize_grid"`
	NvimUnsubscribe         APIUnsubscribe         `json:"nvim_unsubscribe" msgpack:"nvim_unsubscribe"`
	NvimWinClose            APIWinClose            `json:"nvim_win_close" msgpack:"nvim_win_close"`
	NvimWinDelVar           APIWinDelVar           `json:"nvim_win_del_var" msgpack:"nvim_win_del_var"`
	NvimWinGetBuf           APIWinGetBuf           `json:"nvim_win_get_buf" msgpack:"nvim_win_get_buf"`
	NvimWinGetConfig        APIWinGetConfig        `json:"nvim_win_get_config" msgpack:"nvim_win_get_config"`
	NvimWinGetCursor        APIWinGetCursor        `json:"nvim_win_get_cursor" msgpack:"nvim_win_get_cursor"`
	NvimWinGetHeight        APIWinGetHeight        `json:"nvim_win_get_height" msgpack:"nvim_win_get_height"`
	NvimWinGetNumber        APIWinGetNumber        `json:"nvim_win_get_number" msgpack:"nvim_win_get_number"`
	NvimWinGetOption        APIWinGetOption        `json:"nvim_win_get_option" msgpack:"nvim_win_get_option"`
	NvimWinGetPosition      APIWinGetPosition      `json:"nvim_win_get_position" msgpack:"nvim_win_get_position"`
	NvimWinGetTabpage       APIWinGetTabpage       `json:"nvim_win_get_tabpage" msgpack:"nvim_win_get_tabpage"`
	NvimWinGetVar           APIWinGetVar           `json:"nvim_win_get_var" msgpack:"nvim_win_get_var"`
	NvimWinGetWidth         APIWinGetWidth         `json:"nvim_win_get_width" msgpack:"nvim_win_get_width"`
	NvimWinIsValid          APIWinIsValid          `json:"nvim_win_is_valid" msgpack:"nvim_win_is_valid"`
	NvimWinSetBuf           APIWinSetBuf           `json:"nvim_win_set_buf" msgpack:"nvim_win_set_buf"`
	NvimWinSetConfig        APIWinSetConfig        `json:"nvim_win_set_config" msgpack:"nvim_win_set_config"`
	NvimWinSetCursor        APIWinSetCursor        `json:"nvim_win_set_cursor" msgpack:"nvim_win_set_cursor"`
	NvimWinSetHeight        APIWinSetHeight        `json:"nvim_win_set_height" msgpack:"nvim_win_set_height"`
	NvimWinSetOption        APIWinSetOption        `json:"nvim_win_set_option" msgpack:"nvim_win_set_option"`
	NvimWinSetVar           APIWinSetVar           `json:"nvim_win_set_var" msgpack:"nvim_win_set_var"`
	NvimWinSetWidth         APIWinSetWidth         `json:"nvim_win_set_width" msgpack:"nvim_win_set_width"`
}

type APIBufAddDecoration struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []interface{} `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []interface{} `json:"return,omitempty" msgpack:"return,omitempty"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIBufRedrawRange struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []interface{} `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []interface{} `json:"return,omitempty" msgpack:"return,omitempty"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIBufSetLuahl struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []interface{} `json:"return,omitempty" msgpack:"return,omitempty"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIBufStats struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []interface{} `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []interface{} `json:"return,omitempty" msgpack:"return,omitempty"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIGetLibDir struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []interface{} `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []interface{} `json:"return,omitempty" msgpack:"return,omitempty"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIID struct {
	Annotations   []interface{}      `json:"annotations" msgpack:"annotations"`
	Doc           []string           `json:"doc" msgpack:"doc"`
	Parameters    []interface{}      `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIIDParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string           `json:"return" msgpack:"return"`
	Seealso       []interface{}      `json:"seealso" msgpack:"seealso"`
	Signature     string             `json:"signature" msgpack:"signature"`
}

type APIIDParametersDoc struct {
	Obj string `json:"obj" msgpack:"obj"`
}

type APIIDArray struct {
	Annotations   []interface{}           `json:"annotations" msgpack:"annotations"`
	Doc           []string                `json:"doc" msgpack:"doc"`
	Parameters    []interface{}           `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIIDArrayParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                `json:"return" msgpack:"return"`
	Seealso       []interface{}           `json:"seealso" msgpack:"seealso"`
	Signature     string                  `json:"signature" msgpack:"signature"`
}

type APIIDArrayParametersDoc struct {
	Arr string `json:"arr" msgpack:"arr"`
}

type APIIDDictionary struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIIDDictionaryParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                     `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIIDDictionaryParametersDoc struct {
	Dct string `json:"dct" msgpack:"dct"`
}

type APIIDFloat struct {
	Annotations   []interface{}           `json:"annotations" msgpack:"annotations"`
	Doc           []string                `json:"doc" msgpack:"doc"`
	Parameters    []interface{}           `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIIDFloatParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                `json:"return" msgpack:"return"`
	Seealso       []interface{}           `json:"seealso" msgpack:"seealso"`
	Signature     string                  `json:"signature" msgpack:"signature"`
}

type APIIDFloatParametersDoc struct {
	Flt string `json:"flt" msgpack:"flt"`
}

type APIInspectCell struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []interface{} `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []interface{} `json:"return,omitempty" msgpack:"return,omitempty"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIPutAttr struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []interface{} `json:"return,omitempty" msgpack:"return,omitempty"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIStats struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIBufAddHighlight struct {
	Annotations   []interface{}                   `json:"annotations" msgpack:"annotations"`
	Doc           []string                        `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                   `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufAddHighlightParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                        `json:"return" msgpack:"return"`
	Seealso       []interface{}                   `json:"seealso" msgpack:"seealso"`
	Signature     string                          `json:"signature" msgpack:"signature"`
}

type APIBufAddHighlightParametersDoc struct {
	Buffer   string `json:"buffer" msgpack:"buffer"`
	ColEnd   string `json:"col_end" msgpack:"col_end"`
	ColStart string `json:"col_start" msgpack:"col_start"`
	HlGroup  string `json:"hl_group" msgpack:"hl_group"`
	Line     string `json:"line" msgpack:"line"`
	NsID     string `json:"ns_id" msgpack:"ns_id"`
}

type APIBufAttach struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufAttachParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                  `json:"return" msgpack:"return"`
	Seealso       []string                  `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APIBufAttachParametersDoc struct {
	Buffer     string `json:"buffer" msgpack:"buffer"`
	Opts       string `json:"opts" msgpack:"opts"`
	SendBuffer string `json:"send_buffer" msgpack:"send_buffer"`
}

type APIBufClearNamespace struct {
	Annotations   []interface{}                     `json:"annotations" msgpack:"annotations"`
	Doc           []string                          `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                     `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufClearNamespaceParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                     `json:"return" msgpack:"return"`
	Seealso       []interface{}                     `json:"seealso" msgpack:"seealso"`
	Signature     string                            `json:"signature" msgpack:"signature"`
}

type APIBufClearNamespaceParametersDoc struct {
	Buffer    string `json:"buffer" msgpack:"buffer"`
	LineEnd   string `json:"line_end" msgpack:"line_end"`
	LineStart string `json:"line_start" msgpack:"line_start"`
	NsID      string `json:"ns_id" msgpack:"ns_id"`
}

type APIBufDelExtmark struct {
	Annotations   []interface{}                 `json:"annotations" msgpack:"annotations"`
	Doc           []string                      `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                 `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufDelExtmarkParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                      `json:"return" msgpack:"return"`
	Seealso       []interface{}                 `json:"seealso" msgpack:"seealso"`
	Signature     string                        `json:"signature" msgpack:"signature"`
}

type APIBufDelExtmarkParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	ID     string `json:"id" msgpack:"id"`
	NsID   string `json:"ns_id" msgpack:"ns_id"`
}

type APIBufDelKeymap struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufDelKeymapParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                `json:"return" msgpack:"return"`
	Seealso       []string                     `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIBufDelKeymapParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
}

type APIBufDelVar struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufDelVarParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}             `json:"return" msgpack:"return"`
	Seealso       []interface{}             `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APIBufDelVarParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Name   string `json:"name" msgpack:"name"`
}

type APIBufDetach struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufDetachParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                  `json:"return" msgpack:"return"`
	Seealso       []string                  `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APIBufDetachParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
}

type APIBufGetChangedtick struct {
	Annotations   []interface{}                     `json:"annotations" msgpack:"annotations"`
	Doc           []string                          `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                     `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufGetChangedtickParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                          `json:"return" msgpack:"return"`
	Seealso       []interface{}                     `json:"seealso" msgpack:"seealso"`
	Signature     string                            `json:"signature" msgpack:"signature"`
}

type APIBufGetChangedtickParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
}

type APIBufGetCommands struct {
	Annotations   []interface{}                  `json:"annotations" msgpack:"annotations"`
	Doc           []string                       `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                  `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufGetCommandsParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                       `json:"return" msgpack:"return"`
	Seealso       []interface{}                  `json:"seealso" msgpack:"seealso"`
	Signature     string                         `json:"signature" msgpack:"signature"`
}

type APIBufGetCommandsParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Opts   string `json:"opts" msgpack:"opts"`
}

type APIBufGetExtmarkByID struct {
	Annotations   []interface{}                     `json:"annotations" msgpack:"annotations"`
	Doc           []string                          `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                     `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufGetExtmarkByIDParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                          `json:"return" msgpack:"return"`
	Seealso       []interface{}                     `json:"seealso" msgpack:"seealso"`
	Signature     string                            `json:"signature" msgpack:"signature"`
}

type APIBufGetExtmarkByIDParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	ID     string `json:"id" msgpack:"id"`
	NsID   string `json:"ns_id" msgpack:"ns_id"`
}

type APIBufGetExtmarks struct {
	Annotations   []interface{}                  `json:"annotations" msgpack:"annotations"`
	Doc           []string                       `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                  `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufGetExtmarksParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                       `json:"return" msgpack:"return"`
	Seealso       []interface{}                  `json:"seealso" msgpack:"seealso"`
	Signature     string                         `json:"signature" msgpack:"signature"`
}

type APIBufGetExtmarksParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	End    string `json:"end" msgpack:"end"`
	NsID   string `json:"ns_id" msgpack:"ns_id"`
	Opts   string `json:"opts" msgpack:"opts"`
	Start  string `json:"start" msgpack:"start"`
}

type APIBufGetKeymap struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufGetKeymapParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                     `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIBufGetKeymapParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Mode   string `json:"mode" msgpack:"mode"`
}

type APIBufGetLines struct {
	Annotations   []interface{}               `json:"annotations" msgpack:"annotations"`
	Doc           []string                    `json:"doc" msgpack:"doc"`
	Parameters    []interface{}               `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufGetLinesParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                    `json:"return" msgpack:"return"`
	Seealso       []interface{}               `json:"seealso" msgpack:"seealso"`
	Signature     string                      `json:"signature" msgpack:"signature"`
}

type APIBufGetLinesParametersDoc struct {
	Buffer         string `json:"buffer" msgpack:"buffer"`
	End            string `json:"end" msgpack:"end"`
	Start          string `json:"start" msgpack:"start"`
	StrictIndexing string `json:"strict_indexing" msgpack:"strict_indexing"`
}

type APIBufGetMark struct {
	Annotations   []interface{}              `json:"annotations" msgpack:"annotations"`
	Doc           []string                   `json:"doc" msgpack:"doc"`
	Parameters    []interface{}              `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufGetMarkParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                   `json:"return" msgpack:"return"`
	Seealso       []interface{}              `json:"seealso" msgpack:"seealso"`
	Signature     string                     `json:"signature" msgpack:"signature"`
}

type APIBufGetMarkParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Name   string `json:"name" msgpack:"name"`
}

type APIBufGetName struct {
	Annotations   []interface{}              `json:"annotations" msgpack:"annotations"`
	Doc           []string                   `json:"doc" msgpack:"doc"`
	Parameters    []interface{}              `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufGetNameParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                   `json:"return" msgpack:"return"`
	Seealso       []interface{}              `json:"seealso" msgpack:"seealso"`
	Signature     string                     `json:"signature" msgpack:"signature"`
}

type APIBufGetNameParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
}

type APIBufGetOffset struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufGetOffsetParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                     `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIBufGetOffsetParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Index  string `json:"index" msgpack:"index"`
}

type APIBufGetOption struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufGetOptionParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                     `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIBufGetOptionParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Name   string `json:"name" msgpack:"name"`
}

type APIBufGetVar struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufGetVarParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                  `json:"return" msgpack:"return"`
	Seealso       []interface{}             `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APIBufGetVarParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Name   string `json:"name" msgpack:"name"`
}

type APIBufGetVirtualText struct {
	Annotations   []interface{}                     `json:"annotations" msgpack:"annotations"`
	Doc           []string                          `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                     `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufGetVirtualTextParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                          `json:"return" msgpack:"return"`
	Seealso       []interface{}                     `json:"seealso" msgpack:"seealso"`
	Signature     string                            `json:"signature" msgpack:"signature"`
}

type APIBufGetVirtualTextParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Line   string `json:"line" msgpack:"line"`
}

type APIBufIsLoaded struct {
	Annotations   []interface{}               `json:"annotations" msgpack:"annotations"`
	Doc           []string                    `json:"doc" msgpack:"doc"`
	Parameters    []interface{}               `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufIsLoadedParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                    `json:"return" msgpack:"return"`
	Seealso       []interface{}               `json:"seealso" msgpack:"seealso"`
	Signature     string                      `json:"signature" msgpack:"signature"`
}

type APIBufIsLoadedParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
}

type APIBufIsValid struct {
	Annotations   []interface{}              `json:"annotations" msgpack:"annotations"`
	Doc           []string                   `json:"doc" msgpack:"doc"`
	Parameters    []interface{}              `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufIsValidParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                   `json:"return" msgpack:"return"`
	Seealso       []interface{}              `json:"seealso" msgpack:"seealso"`
	Signature     string                     `json:"signature" msgpack:"signature"`
}

type APIBufIsValidParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
}

type APIBufLineCount struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufLineCountParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                     `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIBufLineCountParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
}

type APIBufSetExtmark struct {
	Annotations   []interface{}                 `json:"annotations" msgpack:"annotations"`
	Doc           []string                      `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                 `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufSetExtmarkParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                      `json:"return" msgpack:"return"`
	Seealso       []interface{}                 `json:"seealso" msgpack:"seealso"`
	Signature     string                        `json:"signature" msgpack:"signature"`
}

type APIBufSetExtmarkParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Col    string `json:"col" msgpack:"col"`
	ID     string `json:"id" msgpack:"id"`
	Line   string `json:"line" msgpack:"line"`
	NsID   string `json:"ns_id" msgpack:"ns_id"`
	Opts   string `json:"opts" msgpack:"opts"`
}

type APIBufSetKeymap struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufSetKeymapParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                `json:"return" msgpack:"return"`
	Seealso       []string                     `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIBufSetKeymapParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
}

type APIBufSetLines struct {
	Annotations   []interface{}               `json:"annotations" msgpack:"annotations"`
	Doc           []string                    `json:"doc" msgpack:"doc"`
	Parameters    []interface{}               `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufSetLinesParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}               `json:"return" msgpack:"return"`
	Seealso       []interface{}               `json:"seealso" msgpack:"seealso"`
	Signature     string                      `json:"signature" msgpack:"signature"`
}

type APIBufSetLinesParametersDoc struct {
	Buffer         string `json:"buffer" msgpack:"buffer"`
	End            string `json:"end" msgpack:"end"`
	Replacement    string `json:"replacement" msgpack:"replacement"`
	Start          string `json:"start" msgpack:"start"`
	StrictIndexing string `json:"strict_indexing" msgpack:"strict_indexing"`
}

type APIBufSetName struct {
	Annotations   []interface{}              `json:"annotations" msgpack:"annotations"`
	Doc           []string                   `json:"doc" msgpack:"doc"`
	Parameters    []interface{}              `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufSetNameParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}              `json:"return" msgpack:"return"`
	Seealso       []interface{}              `json:"seealso" msgpack:"seealso"`
	Signature     string                     `json:"signature" msgpack:"signature"`
}

type APIBufSetNameParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Name   string `json:"name" msgpack:"name"`
}

type APIBufSetOption struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufSetOptionParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIBufSetOptionParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Name   string `json:"name" msgpack:"name"`
	Value  string `json:"value" msgpack:"value"`
}

type APIBufSetVar struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufSetVarParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}             `json:"return" msgpack:"return"`
	Seealso       []interface{}             `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APIBufSetVarParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Name   string `json:"name" msgpack:"name"`
	Value  string `json:"value" msgpack:"value"`
}

type APIBufSetVirtualText struct {
	Annotations   []interface{}                     `json:"annotations" msgpack:"annotations"`
	Doc           []string                          `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                     `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIBufSetVirtualTextParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                          `json:"return" msgpack:"return"`
	Seealso       []interface{}                     `json:"seealso" msgpack:"seealso"`
	Signature     string                            `json:"signature" msgpack:"signature"`
}

type APIBufSetVirtualTextParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Chunks string `json:"chunks" msgpack:"chunks"`
	Line   string `json:"line" msgpack:"line"`
	NsID   string `json:"ns_id" msgpack:"ns_id"`
	Opts   string `json:"opts" msgpack:"opts"`
}

type APICallAtomic struct {
	Annotations   []interface{}              `json:"annotations" msgpack:"annotations"`
	Doc           []string                   `json:"doc" msgpack:"doc"`
	Parameters    []interface{}              `json:"parameters" msgpack:"parameters"`
	ParametersDoc APICallAtomicParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                   `json:"return" msgpack:"return"`
	Seealso       []interface{}              `json:"seealso" msgpack:"seealso"`
	Signature     string                     `json:"signature" msgpack:"signature"`
}

type APICallAtomicParametersDoc struct {
	Calls string `json:"calls" msgpack:"calls"`
}

type APICallDictFunction struct {
	Annotations   []interface{}                    `json:"annotations" msgpack:"annotations"`
	Doc           []string                         `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                    `json:"parameters" msgpack:"parameters"`
	ParametersDoc APICallDictFunctionParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                         `json:"return" msgpack:"return"`
	Seealso       []interface{}                    `json:"seealso" msgpack:"seealso"`
	Signature     string                           `json:"signature" msgpack:"signature"`
}

type APICallDictFunctionParametersDoc struct {
	Args string `json:"args" msgpack:"args"`
	Dict string `json:"dict" msgpack:"dict"`
	Fn   string `json:"fn" msgpack:"fn"`
}

type APICallFunction struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APICallFunctionParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                     `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APICallFunctionParametersDoc struct {
	Args string `json:"args" msgpack:"args"`
	Fn   string `json:"fn" msgpack:"fn"`
}

type APICommand struct {
	Annotations   []interface{}           `json:"annotations" msgpack:"annotations"`
	Doc           []string                `json:"doc" msgpack:"doc"`
	Parameters    []interface{}           `json:"parameters" msgpack:"parameters"`
	ParametersDoc APICommandParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}           `json:"return" msgpack:"return"`
	Seealso       []string                `json:"seealso" msgpack:"seealso"`
	Signature     string                  `json:"signature" msgpack:"signature"`
}

type APICommandParametersDoc struct {
	Command string `json:"command" msgpack:"command"`
}

type APICreateBuf struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APICreateBufParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                  `json:"return" msgpack:"return"`
	Seealso       []string                  `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APICreateBufParametersDoc struct {
	Listed  string `json:"listed" msgpack:"listed"`
	Scratch string `json:"scratch" msgpack:"scratch"`
}

type APICreateNamespace struct {
	Annotations   []interface{}                   `json:"annotations" msgpack:"annotations"`
	Doc           []string                        `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                   `json:"parameters" msgpack:"parameters"`
	ParametersDoc APICreateNamespaceParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                        `json:"return" msgpack:"return"`
	Seealso       []interface{}                   `json:"seealso" msgpack:"seealso"`
	Signature     string                          `json:"signature" msgpack:"signature"`
}

type APICreateNamespaceParametersDoc struct {
	Name string `json:"name" msgpack:"name"`
}

type APIDelCurrentLine struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []interface{} `json:"return,omitempty" msgpack:"return,omitempty"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIDelKeymap struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []interface{} `json:"return,omitempty" msgpack:"return,omitempty"`
	Seealso       []string      `json:"seealso" msgpack:"seealso"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIDelVar struct {
	Annotations   []interface{}          `json:"annotations" msgpack:"annotations"`
	Doc           []string               `json:"doc" msgpack:"doc"`
	Parameters    []interface{}          `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIDelVarParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}          `json:"return" msgpack:"return"`
	Seealso       []interface{}          `json:"seealso" msgpack:"seealso"`
	Signature     string                 `json:"signature" msgpack:"signature"`
}

type APIDelVarParametersDoc struct {
	Name string `json:"name" msgpack:"name"`
}

type APIErrWrite struct {
	Annotations   []interface{}            `json:"annotations" msgpack:"annotations"`
	Doc           []string                 `json:"doc" msgpack:"doc"`
	Parameters    []interface{}            `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIErrWriteParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}            `json:"return" msgpack:"return"`
	Seealso       []interface{}            `json:"seealso" msgpack:"seealso"`
	Signature     string                   `json:"signature" msgpack:"signature"`
}

type APIErrWriteParametersDoc struct {
	Str string `json:"str" msgpack:"str"`
}

type APIErrWriteln struct {
	Annotations   []interface{}              `json:"annotations" msgpack:"annotations"`
	Doc           []string                   `json:"doc" msgpack:"doc"`
	Parameters    []interface{}              `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIErrWritelnParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}              `json:"return" msgpack:"return"`
	Seealso       []string                   `json:"seealso" msgpack:"seealso"`
	Signature     string                     `json:"signature" msgpack:"signature"`
}

type APIErrWritelnParametersDoc struct {
	Str string `json:"str" msgpack:"str"`
}

type APIEval struct {
	Annotations   []interface{}        `json:"annotations" msgpack:"annotations"`
	Doc           []string             `json:"doc" msgpack:"doc"`
	Parameters    []interface{}        `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIEvalParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string             `json:"return" msgpack:"return"`
	Seealso       []interface{}        `json:"seealso" msgpack:"seealso"`
	Signature     string               `json:"signature" msgpack:"signature"`
}

type APIEvalParametersDoc struct {
	Expr string `json:"expr" msgpack:"expr"`
}

type APIExec struct {
	Annotations   []interface{}        `json:"annotations" msgpack:"annotations"`
	Doc           []string             `json:"doc" msgpack:"doc"`
	Parameters    []interface{}        `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIExecParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string             `json:"return" msgpack:"return"`
	Seealso       []string             `json:"seealso" msgpack:"seealso"`
	Signature     string               `json:"signature" msgpack:"signature"`
}

type APIExecParametersDoc struct {
	Output string `json:"output" msgpack:"output"`
	Src    string `json:"src" msgpack:"src"`
}

type APIExecLua struct {
	Annotations   []interface{}           `json:"annotations" msgpack:"annotations"`
	Doc           []string                `json:"doc" msgpack:"doc"`
	Parameters    []interface{}           `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIExecLuaParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                `json:"return" msgpack:"return"`
	Seealso       []interface{}           `json:"seealso" msgpack:"seealso"`
	Signature     string                  `json:"signature" msgpack:"signature"`
}

type APIExecLuaParametersDoc struct {
	Args string `json:"args" msgpack:"args"`
	Code string `json:"code" msgpack:"code"`
}

type APIFeedkeys struct {
	Annotations   []interface{}            `json:"annotations" msgpack:"annotations"`
	Doc           []string                 `json:"doc" msgpack:"doc"`
	Parameters    []interface{}            `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIFeedkeysParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}            `json:"return" msgpack:"return"`
	Seealso       []string                 `json:"seealso" msgpack:"seealso"`
	Signature     string                   `json:"signature" msgpack:"signature"`
}

type APIFeedkeysParametersDoc struct {
	EscapeCsi string `json:"escape_csi" msgpack:"escape_csi"`
	Keys      string `json:"keys" msgpack:"keys"`
	Mode      string `json:"mode" msgpack:"mode"`
}

type APIGetAPIInfo struct {
	Annotations   []string      `json:"annotations" msgpack:"annotations"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIGetChanInfo struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIGetColorByName struct {
	Annotations   []interface{}                  `json:"annotations" msgpack:"annotations"`
	Doc           []string                       `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                  `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIGetColorByNameParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                       `json:"return" msgpack:"return"`
	Seealso       []interface{}                  `json:"seealso" msgpack:"seealso"`
	Signature     string                         `json:"signature" msgpack:"signature"`
}

type APIGetColorByNameParametersDoc struct {
	Name string `json:"name" msgpack:"name"`
}

type APIGetColorMap struct {
	Annotations   []interface{}               `json:"annotations" msgpack:"annotations"`
	Doc           []string                    `json:"doc" msgpack:"doc"`
	Parameters    []interface{}               `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIGetColorMapParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                    `json:"return" msgpack:"return"`
	Seealso       []interface{}               `json:"seealso" msgpack:"seealso"`
	Signature     string                      `json:"signature" msgpack:"signature"`
}

type APIGetColorMapParametersDoc struct {
}

type APIGetCommands struct {
	Annotations   []interface{}               `json:"annotations" msgpack:"annotations"`
	Doc           []string                    `json:"doc" msgpack:"doc"`
	Parameters    []interface{}               `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIGetCommandsParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                    `json:"return" msgpack:"return"`
	Seealso       []interface{}               `json:"seealso" msgpack:"seealso"`
	Signature     string                      `json:"signature" msgpack:"signature"`
}

type APIGetCommandsParametersDoc struct {
	Opts string `json:"opts" msgpack:"opts"`
}

type APIGetContext struct {
	Annotations   []interface{}              `json:"annotations" msgpack:"annotations"`
	Doc           []string                   `json:"doc" msgpack:"doc"`
	Parameters    []interface{}              `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIGetContextParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                   `json:"return" msgpack:"return"`
	Seealso       []interface{}              `json:"seealso" msgpack:"seealso"`
	Signature     string                     `json:"signature" msgpack:"signature"`
}

type APIGetContextParametersDoc struct {
	Opts string `json:"opts" msgpack:"opts"`
}

type APIGetCurrentBuf struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIGetCurrentLine struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIGetCurrentTabpage struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIGetCurrentWin struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIGetHlByID struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIGetHlByIDParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                  `json:"return" msgpack:"return"`
	Seealso       []string                  `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APIGetHlByIDParametersDoc struct {
	HlID string `json:"hl_id" msgpack:"hl_id"`
	Rgb  string `json:"rgb" msgpack:"rgb"`
}

type APIGetHlByName struct {
	Annotations   []interface{}               `json:"annotations" msgpack:"annotations"`
	Doc           []string                    `json:"doc" msgpack:"doc"`
	Parameters    []interface{}               `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIGetHlByNameParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                    `json:"return" msgpack:"return"`
	Seealso       []string                    `json:"seealso" msgpack:"seealso"`
	Signature     string                      `json:"signature" msgpack:"signature"`
}

type APIGetHlByNameParametersDoc struct {
	Name string `json:"name" msgpack:"name"`
	Rgb  string `json:"rgb" msgpack:"rgb"`
}

type APIGetHlIDByName struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []interface{} `json:"return,omitempty" msgpack:"return,omitempty"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIGetKeymap struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIGetKeymapParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                  `json:"return" msgpack:"return"`
	Seealso       []interface{}             `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APIGetKeymapParametersDoc struct {
	Mode string `json:"mode" msgpack:"mode"`
}

type APIGetMode struct {
	Annotations   []string      `json:"annotations" msgpack:"annotations"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIGetNamespaces struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIGetOption struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIGetOptionParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                  `json:"return" msgpack:"return"`
	Seealso       []interface{}             `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APIGetOptionParametersDoc struct {
	Name string `json:"name" msgpack:"name"`
}

type APIGetProc struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIGetProcChildren struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIGetRuntimeFile struct {
	Annotations   []interface{}                  `json:"annotations" msgpack:"annotations"`
	Doc           []string                       `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                  `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIGetRuntimeFileParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                       `json:"return" msgpack:"return"`
	Seealso       []interface{}                  `json:"seealso" msgpack:"seealso"`
	Signature     string                         `json:"signature" msgpack:"signature"`
}

type APIGetRuntimeFileParametersDoc struct {
	All  string `json:"all" msgpack:"all"`
	Name string `json:"name" msgpack:"name"`
}

type APIGetVar struct {
	Annotations   []interface{}          `json:"annotations" msgpack:"annotations"`
	Doc           []string               `json:"doc" msgpack:"doc"`
	Parameters    []interface{}          `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIGetVarParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string               `json:"return" msgpack:"return"`
	Seealso       []interface{}          `json:"seealso" msgpack:"seealso"`
	Signature     string                 `json:"signature" msgpack:"signature"`
}

type APIGetVarParametersDoc struct {
	Name string `json:"name" msgpack:"name"`
}

type APIGetVvar struct {
	Annotations   []interface{}           `json:"annotations" msgpack:"annotations"`
	Doc           []string                `json:"doc" msgpack:"doc"`
	Parameters    []interface{}           `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIGetVvarParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                `json:"return" msgpack:"return"`
	Seealso       []interface{}           `json:"seealso" msgpack:"seealso"`
	Signature     string                  `json:"signature" msgpack:"signature"`
}

type APIGetVvarParametersDoc struct {
	Name string `json:"name" msgpack:"name"`
}

type APIInput struct {
	Annotations   []string              `json:"annotations" msgpack:"annotations"`
	Doc           []string              `json:"doc" msgpack:"doc"`
	Parameters    []interface{}         `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIInputParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string              `json:"return" msgpack:"return"`
	Seealso       []interface{}         `json:"seealso" msgpack:"seealso"`
	Signature     string                `json:"signature" msgpack:"signature"`
}

type APIInputParametersDoc struct {
	Keys string `json:"keys" msgpack:"keys"`
}

type APIInputMouse struct {
	Annotations   []string                   `json:"annotations" msgpack:"annotations"`
	Doc           []string                   `json:"doc" msgpack:"doc"`
	Parameters    []interface{}              `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIInputMouseParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}              `json:"return" msgpack:"return"`
	Seealso       []interface{}              `json:"seealso" msgpack:"seealso"`
	Signature     string                     `json:"signature" msgpack:"signature"`
}

type APIInputMouseParametersDoc struct {
	Action   string `json:"action" msgpack:"action"`
	Button   string `json:"button" msgpack:"button"`
	Col      string `json:"col" msgpack:"col"`
	Grid     string `json:"grid" msgpack:"grid"`
	Modifier string `json:"modifier" msgpack:"modifier"`
	Row      string `json:"row" msgpack:"row"`
}

type APIListBufs struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIListChans struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIListRuntimePaths struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIListTabpages struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIListUis struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIListWins struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []string      `json:"return" msgpack:"return"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APILoadContext struct {
	Annotations   []interface{}               `json:"annotations" msgpack:"annotations"`
	Doc           []string                    `json:"doc" msgpack:"doc"`
	Parameters    []interface{}               `json:"parameters" msgpack:"parameters"`
	ParametersDoc APILoadContextParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}               `json:"return" msgpack:"return"`
	Seealso       []interface{}               `json:"seealso" msgpack:"seealso"`
	Signature     string                      `json:"signature" msgpack:"signature"`
}

type APILoadContextParametersDoc struct {
	Dict string `json:"dict" msgpack:"dict"`
}

type APIOpenWin struct {
	Annotations   []interface{}           `json:"annotations" msgpack:"annotations"`
	Doc           []string                `json:"doc" msgpack:"doc"`
	Parameters    []interface{}           `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIOpenWinParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                `json:"return" msgpack:"return"`
	Seealso       []interface{}           `json:"seealso" msgpack:"seealso"`
	Signature     string                  `json:"signature" msgpack:"signature"`
}

type APIOpenWinParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Config string `json:"config" msgpack:"config"`
	Enter  string `json:"enter" msgpack:"enter"`
}

type APIOutWrite struct {
	Annotations   []interface{}            `json:"annotations" msgpack:"annotations"`
	Doc           []string                 `json:"doc" msgpack:"doc"`
	Parameters    []interface{}            `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIOutWriteParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}            `json:"return" msgpack:"return"`
	Seealso       []interface{}            `json:"seealso" msgpack:"seealso"`
	Signature     string                   `json:"signature" msgpack:"signature"`
}

type APIOutWriteParametersDoc struct {
	Str string `json:"str" msgpack:"str"`
}

type APIParseExpression struct {
	Annotations   []string                        `json:"annotations" msgpack:"annotations"`
	Doc           []string                        `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                   `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIParseExpressionParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                        `json:"return" msgpack:"return"`
	Seealso       []interface{}                   `json:"seealso" msgpack:"seealso"`
	Signature     string                          `json:"signature" msgpack:"signature"`
}

type APIParseExpressionParametersDoc struct {
	Expr      string `json:"expr" msgpack:"expr"`
	Flags     string `json:"flags" msgpack:"flags"`
	Highlight string `json:"highlight" msgpack:"highlight"`
}

type APIPaste struct {
	Annotations   []interface{}         `json:"annotations" msgpack:"annotations"`
	Doc           []string              `json:"doc" msgpack:"doc"`
	Parameters    []interface{}         `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIPasteParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string              `json:"return" msgpack:"return"`
	Seealso       []interface{}         `json:"seealso" msgpack:"seealso"`
	Signature     string                `json:"signature" msgpack:"signature"`
}

type APIPasteParametersDoc struct {
	Crlf  string `json:"crlf" msgpack:"crlf"`
	Data  string `json:"data" msgpack:"data"`
	Phase string `json:"phase" msgpack:"phase"`
}

type APIPut struct {
	Annotations   []interface{}       `json:"annotations" msgpack:"annotations"`
	Doc           []string            `json:"doc" msgpack:"doc"`
	Parameters    []interface{}       `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIPutParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}       `json:"return" msgpack:"return"`
	Seealso       []interface{}       `json:"seealso" msgpack:"seealso"`
	Signature     string              `json:"signature" msgpack:"signature"`
}

type APIPutParametersDoc struct {
	After  string `json:"after" msgpack:"after"`
	Follow string `json:"follow" msgpack:"follow"`
	Lines  string `json:"lines" msgpack:"lines"`
	Type   string `json:"type" msgpack:"type"`
}

type APIReplaceTermcodes struct {
	Annotations   []interface{}                    `json:"annotations" msgpack:"annotations"`
	Doc           []string                         `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                    `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIReplaceTermcodesParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                    `json:"return" msgpack:"return"`
	Seealso       []string                         `json:"seealso" msgpack:"seealso"`
	Signature     string                           `json:"signature" msgpack:"signature"`
}

type APIReplaceTermcodesParametersDoc struct {
	DoLt     string `json:"do_lt" msgpack:"do_lt"`
	FromPart string `json:"from_part" msgpack:"from_part"`
	Special  string `json:"special" msgpack:"special"`
	Str      string `json:"str" msgpack:"str"`
}

type APISelectPopupmenuItem struct {
	Annotations   []interface{}                       `json:"annotations" msgpack:"annotations"`
	Doc           []string                            `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                       `json:"parameters" msgpack:"parameters"`
	ParametersDoc APISelectPopupmenuItemParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                       `json:"return" msgpack:"return"`
	Seealso       []interface{}                       `json:"seealso" msgpack:"seealso"`
	Signature     string                              `json:"signature" msgpack:"signature"`
}

type APISelectPopupmenuItemParametersDoc struct {
	Finish string `json:"finish" msgpack:"finish"`
	Insert string `json:"insert" msgpack:"insert"`
	Item   string `json:"item" msgpack:"item"`
	Opts   string `json:"opts" msgpack:"opts"`
}

type APISetClientInfo struct {
	Annotations   []interface{}                 `json:"annotations" msgpack:"annotations"`
	Doc           []string                      `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                 `json:"parameters" msgpack:"parameters"`
	ParametersDoc APISetClientInfoParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                 `json:"return" msgpack:"return"`
	Seealso       []interface{}                 `json:"seealso" msgpack:"seealso"`
	Signature     string                        `json:"signature" msgpack:"signature"`
}

type APISetClientInfoParametersDoc struct {
	Attributes string `json:"attributes" msgpack:"attributes"`
	Methods    string `json:"methods" msgpack:"methods"`
	Name       string `json:"name" msgpack:"name"`
	Type       string `json:"type" msgpack:"type"`
	Version    string `json:"version" msgpack:"version"`
}

type APISetCurrentBuf struct {
	Annotations   []interface{}                 `json:"annotations" msgpack:"annotations"`
	Doc           []string                      `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                 `json:"parameters" msgpack:"parameters"`
	ParametersDoc APISetCurrentBufParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                 `json:"return" msgpack:"return"`
	Seealso       []interface{}                 `json:"seealso" msgpack:"seealso"`
	Signature     string                        `json:"signature" msgpack:"signature"`
}

type APISetCurrentBufParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
}

type APISetCurrentDir struct {
	Annotations   []interface{}                 `json:"annotations" msgpack:"annotations"`
	Doc           []string                      `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                 `json:"parameters" msgpack:"parameters"`
	ParametersDoc APISetCurrentDirParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                 `json:"return" msgpack:"return"`
	Seealso       []interface{}                 `json:"seealso" msgpack:"seealso"`
	Signature     string                        `json:"signature" msgpack:"signature"`
}

type APISetCurrentDirParametersDoc struct {
	Dir string `json:"dir" msgpack:"dir"`
}

type APISetCurrentLine struct {
	Annotations   []interface{}                  `json:"annotations" msgpack:"annotations"`
	Doc           []string                       `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                  `json:"parameters" msgpack:"parameters"`
	ParametersDoc APISetCurrentLineParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                  `json:"return" msgpack:"return"`
	Seealso       []interface{}                  `json:"seealso" msgpack:"seealso"`
	Signature     string                         `json:"signature" msgpack:"signature"`
}

type APISetCurrentLineParametersDoc struct {
	Line string `json:"line" msgpack:"line"`
}

type APISetCurrentTabpage struct {
	Annotations   []interface{}                     `json:"annotations" msgpack:"annotations"`
	Doc           []string                          `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                     `json:"parameters" msgpack:"parameters"`
	ParametersDoc APISetCurrentTabpageParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                     `json:"return" msgpack:"return"`
	Seealso       []interface{}                     `json:"seealso" msgpack:"seealso"`
	Signature     string                            `json:"signature" msgpack:"signature"`
}

type APISetCurrentTabpageParametersDoc struct {
	Tabpage string `json:"tabpage" msgpack:"tabpage"`
}

type APISetCurrentWin struct {
	Annotations   []interface{}                 `json:"annotations" msgpack:"annotations"`
	Doc           []string                      `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                 `json:"parameters" msgpack:"parameters"`
	ParametersDoc APISetCurrentWinParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                 `json:"return" msgpack:"return"`
	Seealso       []interface{}                 `json:"seealso" msgpack:"seealso"`
	Signature     string                        `json:"signature" msgpack:"signature"`
}

type APISetCurrentWinParametersDoc struct {
	Window string `json:"window" msgpack:"window"`
}

type APISetKeymap struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APISetKeymapParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}             `json:"return" msgpack:"return"`
	Seealso       []interface{}             `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APISetKeymapParametersDoc struct {
	LHS  string `json:"lhs" msgpack:"lhs"`
	Mode string `json:"mode" msgpack:"mode"`
	Opts string `json:"opts" msgpack:"opts"`
	RHS  string `json:"rhs" msgpack:"rhs"`
}

type APISetOption struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APISetOptionParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}             `json:"return" msgpack:"return"`
	Seealso       []interface{}             `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APISetOptionParametersDoc struct {
	Name  string `json:"name" msgpack:"name"`
	Value string `json:"value" msgpack:"value"`
}

type APISetVar struct {
	Annotations   []interface{}          `json:"annotations" msgpack:"annotations"`
	Doc           []string               `json:"doc" msgpack:"doc"`
	Parameters    []interface{}          `json:"parameters" msgpack:"parameters"`
	ParametersDoc APISetVarParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}          `json:"return" msgpack:"return"`
	Seealso       []interface{}          `json:"seealso" msgpack:"seealso"`
	Signature     string                 `json:"signature" msgpack:"signature"`
}

type APISetVarParametersDoc struct {
	Name  string `json:"name" msgpack:"name"`
	Value string `json:"value" msgpack:"value"`
}

type APISetVvar struct {
	Annotations   []interface{}           `json:"annotations" msgpack:"annotations"`
	Doc           []string                `json:"doc" msgpack:"doc"`
	Parameters    []interface{}           `json:"parameters" msgpack:"parameters"`
	ParametersDoc APISetVvarParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}           `json:"return" msgpack:"return"`
	Seealso       []interface{}           `json:"seealso" msgpack:"seealso"`
	Signature     string                  `json:"signature" msgpack:"signature"`
}

type APISetVvarParametersDoc struct {
	Name  string `json:"name" msgpack:"name"`
	Value string `json:"value" msgpack:"value"`
}

type APIStrwidth struct {
	Annotations   []interface{}            `json:"annotations" msgpack:"annotations"`
	Doc           []string                 `json:"doc" msgpack:"doc"`
	Parameters    []interface{}            `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIStrwidthParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                 `json:"return" msgpack:"return"`
	Seealso       []interface{}            `json:"seealso" msgpack:"seealso"`
	Signature     string                   `json:"signature" msgpack:"signature"`
}

type APIStrwidthParametersDoc struct {
	Text string `json:"text" msgpack:"text"`
}

type APISubscribe struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APISubscribeParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}             `json:"return" msgpack:"return"`
	Seealso       []interface{}             `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APISubscribeParametersDoc struct {
	Event string `json:"event" msgpack:"event"`
}

type APITabpageDelVar struct {
	Annotations   []interface{}                 `json:"annotations" msgpack:"annotations"`
	Doc           []string                      `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                 `json:"parameters" msgpack:"parameters"`
	ParametersDoc APITabpageDelVarParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                 `json:"return" msgpack:"return"`
	Seealso       []interface{}                 `json:"seealso" msgpack:"seealso"`
	Signature     string                        `json:"signature" msgpack:"signature"`
}

type APITabpageDelVarParametersDoc struct {
	Name    string `json:"name" msgpack:"name"`
	Tabpage string `json:"tabpage" msgpack:"tabpage"`
}

type APITabpageGetNumber struct {
	Annotations   []interface{}                    `json:"annotations" msgpack:"annotations"`
	Doc           []string                         `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                    `json:"parameters" msgpack:"parameters"`
	ParametersDoc APITabpageGetNumberParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                         `json:"return" msgpack:"return"`
	Seealso       []interface{}                    `json:"seealso" msgpack:"seealso"`
	Signature     string                           `json:"signature" msgpack:"signature"`
}

type APITabpageGetNumberParametersDoc struct {
	Tabpage string `json:"tabpage" msgpack:"tabpage"`
}

type APITabpageGetVar struct {
	Annotations   []interface{}                 `json:"annotations" msgpack:"annotations"`
	Doc           []string                      `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                 `json:"parameters" msgpack:"parameters"`
	ParametersDoc APITabpageGetVarParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                      `json:"return" msgpack:"return"`
	Seealso       []interface{}                 `json:"seealso" msgpack:"seealso"`
	Signature     string                        `json:"signature" msgpack:"signature"`
}

type APITabpageGetVarParametersDoc struct {
	Name    string `json:"name" msgpack:"name"`
	Tabpage string `json:"tabpage" msgpack:"tabpage"`
}

type APITabpageGetWin struct {
	Annotations   []interface{}                 `json:"annotations" msgpack:"annotations"`
	Doc           []string                      `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                 `json:"parameters" msgpack:"parameters"`
	ParametersDoc APITabpageGetWinParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                      `json:"return" msgpack:"return"`
	Seealso       []interface{}                 `json:"seealso" msgpack:"seealso"`
	Signature     string                        `json:"signature" msgpack:"signature"`
}

type APITabpageGetWinParametersDoc struct {
	Tabpage string `json:"tabpage" msgpack:"tabpage"`
}

type APITabpageIsValid struct {
	Annotations   []interface{}                  `json:"annotations" msgpack:"annotations"`
	Doc           []string                       `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                  `json:"parameters" msgpack:"parameters"`
	ParametersDoc APITabpageIsValidParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                       `json:"return" msgpack:"return"`
	Seealso       []interface{}                  `json:"seealso" msgpack:"seealso"`
	Signature     string                         `json:"signature" msgpack:"signature"`
}

type APITabpageIsValidParametersDoc struct {
	Tabpage string `json:"tabpage" msgpack:"tabpage"`
}

type APITabpageListWins struct {
	Annotations   []interface{}                   `json:"annotations" msgpack:"annotations"`
	Doc           []string                        `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                   `json:"parameters" msgpack:"parameters"`
	ParametersDoc APITabpageListWinsParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                        `json:"return" msgpack:"return"`
	Seealso       []interface{}                   `json:"seealso" msgpack:"seealso"`
	Signature     string                          `json:"signature" msgpack:"signature"`
}

type APITabpageListWinsParametersDoc struct {
	Tabpage string `json:"tabpage" msgpack:"tabpage"`
}

type APITabpageSetVar struct {
	Annotations   []interface{}                 `json:"annotations" msgpack:"annotations"`
	Doc           []string                      `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                 `json:"parameters" msgpack:"parameters"`
	ParametersDoc APITabpageSetVarParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                 `json:"return" msgpack:"return"`
	Seealso       []interface{}                 `json:"seealso" msgpack:"seealso"`
	Signature     string                        `json:"signature" msgpack:"signature"`
}

type APITabpageSetVarParametersDoc struct {
	Name    string `json:"name" msgpack:"name"`
	Tabpage string `json:"tabpage" msgpack:"tabpage"`
	Value   string `json:"value" msgpack:"value"`
}

type APIUIAttach struct {
	Annotations   []interface{}            `json:"annotations" msgpack:"annotations"`
	Doc           []string                 `json:"doc" msgpack:"doc"`
	Parameters    []interface{}            `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIUIAttachParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}            `json:"return" msgpack:"return"`
	Seealso       []interface{}            `json:"seealso" msgpack:"seealso"`
	Signature     string                   `json:"signature" msgpack:"signature"`
}

type APIUIAttachParametersDoc struct {
	Height  string `json:"height" msgpack:"height"`
	Options string `json:"options" msgpack:"options"`
	Width   string `json:"width" msgpack:"width"`
}

type APIUidetach struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []string      `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []interface{} `json:"return,omitempty" msgpack:"return,omitempty"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIUipumSetHeight struct {
	Annotations   []interface{}                  `json:"annotations" msgpack:"annotations"`
	Doc           []string                       `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                  `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIUipumSetHeightParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                  `json:"return" msgpack:"return"`
	Seealso       []interface{}                  `json:"seealso" msgpack:"seealso"`
	Signature     string                         `json:"signature" msgpack:"signature"`
}

type APIUipumSetHeightParametersDoc struct {
	Height string `json:"height" msgpack:"height"`
}

type APIUISetOption struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []interface{} `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []interface{} `json:"return,omitempty" msgpack:"return,omitempty"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIUITryResize struct {
	Annotations   []interface{} `json:"annotations,omitempty" msgpack:"annotations,omitempty"`
	Doc           []interface{} `json:"doc" msgpack:"doc"`
	Parameters    []interface{} `json:"parameters" msgpack:"parameters"`
	ParametersDoc interface{}   `json:"parameters_doc,omitempty" msgpack:"parameters_doc,omitempty"`
	Return        []interface{} `json:"return,omitempty" msgpack:"return,omitempty"`
	Seealso       []interface{} `json:"seealso,omitempty" msgpack:"seealso,omitempty"`
	Signature     string        `json:"signature" msgpack:"signature"`
}

type APIUITryResizeGrid struct {
	Annotations   []interface{}                   `json:"annotations" msgpack:"annotations"`
	Doc           []string                        `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                   `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIUITryResizeGridParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                   `json:"return" msgpack:"return"`
	Seealso       []interface{}                   `json:"seealso" msgpack:"seealso"`
	Signature     string                          `json:"signature" msgpack:"signature"`
}

type APIUITryResizeGridParametersDoc struct {
	Grid   string `json:"grid" msgpack:"grid"`
	Height string `json:"height" msgpack:"height"`
	Width  string `json:"width" msgpack:"width"`
}

type APIUnsubscribe struct {
	Annotations   []interface{}               `json:"annotations" msgpack:"annotations"`
	Doc           []string                    `json:"doc" msgpack:"doc"`
	Parameters    []interface{}               `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIUnsubscribeParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}               `json:"return" msgpack:"return"`
	Seealso       []interface{}               `json:"seealso" msgpack:"seealso"`
	Signature     string                      `json:"signature" msgpack:"signature"`
}

type APIUnsubscribeParametersDoc struct {
	Event string `json:"event" msgpack:"event"`
}

type APIWinClose struct {
	Annotations   []interface{}            `json:"annotations" msgpack:"annotations"`
	Doc           []string                 `json:"doc" msgpack:"doc"`
	Parameters    []interface{}            `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinCloseParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}            `json:"return" msgpack:"return"`
	Seealso       []interface{}            `json:"seealso" msgpack:"seealso"`
	Signature     string                   `json:"signature" msgpack:"signature"`
}

type APIWinCloseParametersDoc struct {
	Force  string `json:"force" msgpack:"force"`
	Window string `json:"window" msgpack:"window"`
}

type APIWinDelVar struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinDelVarParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}             `json:"return" msgpack:"return"`
	Seealso       []interface{}             `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APIWinDelVarParametersDoc struct {
	Name   string `json:"name" msgpack:"name"`
	Window string `json:"window" msgpack:"window"`
}

type APIWinGetBuf struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinGetBufParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                  `json:"return" msgpack:"return"`
	Seealso       []interface{}             `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APIWinGetBufParametersDoc struct {
	Window string `json:"window" msgpack:"window"`
}

type APIWinGetConfig struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinGetConfigParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                     `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIWinGetConfigParametersDoc struct {
	Window string `json:"window" msgpack:"window"`
}

type APIWinGetCursor struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinGetCursorParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                     `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIWinGetCursorParametersDoc struct {
	Window string `json:"window" msgpack:"window"`
}

type APIWinGetHeight struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinGetHeightParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                     `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIWinGetHeightParametersDoc struct {
	Window string `json:"window" msgpack:"window"`
}

type APIWinGetNumber struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinGetNumberParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                     `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIWinGetNumberParametersDoc struct {
	Window string `json:"window" msgpack:"window"`
}

type APIWinGetOption struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinGetOptionParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                     `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIWinGetOptionParametersDoc struct {
	Name   string `json:"name" msgpack:"name"`
	Window string `json:"window" msgpack:"window"`
}

type APIWinGetPosition struct {
	Annotations   []interface{}                  `json:"annotations" msgpack:"annotations"`
	Doc           []string                       `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                  `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinGetPositionParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                       `json:"return" msgpack:"return"`
	Seealso       []interface{}                  `json:"seealso" msgpack:"seealso"`
	Signature     string                         `json:"signature" msgpack:"signature"`
}

type APIWinGetPositionParametersDoc struct {
	Window string `json:"window" msgpack:"window"`
}

type APIWinGetTabpage struct {
	Annotations   []interface{}                 `json:"annotations" msgpack:"annotations"`
	Doc           []string                      `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                 `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinGetTabpageParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                      `json:"return" msgpack:"return"`
	Seealso       []interface{}                 `json:"seealso" msgpack:"seealso"`
	Signature     string                        `json:"signature" msgpack:"signature"`
}

type APIWinGetTabpageParametersDoc struct {
	Window string `json:"window" msgpack:"window"`
}

type APIWinGetVar struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinGetVarParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                  `json:"return" msgpack:"return"`
	Seealso       []interface{}             `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APIWinGetVarParametersDoc struct {
	Name   string `json:"name" msgpack:"name"`
	Window string `json:"window" msgpack:"window"`
}

type APIWinGetWidth struct {
	Annotations   []interface{}               `json:"annotations" msgpack:"annotations"`
	Doc           []string                    `json:"doc" msgpack:"doc"`
	Parameters    []interface{}               `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinGetWidthParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                    `json:"return" msgpack:"return"`
	Seealso       []interface{}               `json:"seealso" msgpack:"seealso"`
	Signature     string                      `json:"signature" msgpack:"signature"`
}

type APIWinGetWidthParametersDoc struct {
	Window string `json:"window" msgpack:"window"`
}

type APIWinIsValid struct {
	Annotations   []interface{}              `json:"annotations" msgpack:"annotations"`
	Doc           []string                   `json:"doc" msgpack:"doc"`
	Parameters    []interface{}              `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinIsValidParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []string                   `json:"return" msgpack:"return"`
	Seealso       []interface{}              `json:"seealso" msgpack:"seealso"`
	Signature     string                     `json:"signature" msgpack:"signature"`
}

type APIWinIsValidParametersDoc struct {
	Window string `json:"window" msgpack:"window"`
}

type APIWinSetBuf struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinSetBufParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}             `json:"return" msgpack:"return"`
	Seealso       []interface{}             `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APIWinSetBufParametersDoc struct {
	Buffer string `json:"buffer" msgpack:"buffer"`
	Window string `json:"window" msgpack:"window"`
}

type APIWinSetConfig struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinSetConfigParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                `json:"return" msgpack:"return"`
	Seealso       []string                     `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIWinSetConfigParametersDoc struct {
	Config string `json:"config" msgpack:"config"`
	Window string `json:"window" msgpack:"window"`
}

type APIWinSetCursor struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinSetCursorParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIWinSetCursorParametersDoc struct {
	Pos    string `json:"pos" msgpack:"pos"`
	Window string `json:"window" msgpack:"window"`
}

type APIWinSetHeight struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinSetHeightParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIWinSetHeightParametersDoc struct {
	Height string `json:"height" msgpack:"height"`
	Window string `json:"window" msgpack:"window"`
}

type APIWinSetOption struct {
	Annotations   []interface{}                `json:"annotations" msgpack:"annotations"`
	Doc           []string                     `json:"doc" msgpack:"doc"`
	Parameters    []interface{}                `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinSetOptionParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}                `json:"return" msgpack:"return"`
	Seealso       []interface{}                `json:"seealso" msgpack:"seealso"`
	Signature     string                       `json:"signature" msgpack:"signature"`
}

type APIWinSetOptionParametersDoc struct {
	Name   string `json:"name" msgpack:"name"`
	Value  string `json:"value" msgpack:"value"`
	Window string `json:"window" msgpack:"window"`
}

type APIWinSetVar struct {
	Annotations   []interface{}             `json:"annotations" msgpack:"annotations"`
	Doc           []string                  `json:"doc" msgpack:"doc"`
	Parameters    []interface{}             `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinSetVarParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}             `json:"return" msgpack:"return"`
	Seealso       []interface{}             `json:"seealso" msgpack:"seealso"`
	Signature     string                    `json:"signature" msgpack:"signature"`
}

type APIWinSetVarParametersDoc struct {
	Name   string `json:"name" msgpack:"name"`
	Value  string `json:"value" msgpack:"value"`
	Window string `json:"window" msgpack:"window"`
}

type APIWinSetWidth struct {
	Annotations   []interface{}               `json:"annotations" msgpack:"annotations"`
	Doc           []string                    `json:"doc" msgpack:"doc"`
	Parameters    []interface{}               `json:"parameters" msgpack:"parameters"`
	ParametersDoc APIWinSetWidthParametersDoc `json:"parameters_doc" msgpack:"parameters_doc"`
	Return        []interface{}               `json:"return" msgpack:"return"`
	Seealso       []interface{}               `json:"seealso" msgpack:"seealso"`
	Signature     string                      `json:"signature" msgpack:"signature"`
}

type APIWinSetWidthParametersDoc struct {
	Width  string `json:"width" msgpack:"width"`
	Window string `json:"window" msgpack:"window"`
}
