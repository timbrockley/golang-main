/*
################################################################################

RPC Class / Interface

Copyright (C) 2019-2023 Tim Brockley

The JavaScript code used in this file is free software: you can
redistribute it and/or modify it under the terms of the GNU
General Public License (GNU GPL) as published by the Free Software
Foundation, either version 3 of the License, or (at your option)
any later version.  The code is distributed WITHOUT ANY WARRANTY;
without even the implied warranty of MERCHANTABILITY or FITNESS
FOR A PARTICULAR PURPOSE.  See the GNU GPL for more details.

As additional permission under GNU GPL version 3 section 7, you
may distribute non-source (e.g., minimized or compacted) forms of
that code without the copy of the GNU GPL normally required by
section 4, provided you include this license notice and a URL
through which recipients can access the Corresponding Source.

################################################################################
*/
//------------------------------------------------------------
export default class RPC
//------------------------------------------------------------
{
	//############################################################
	//------------------------------------------------------------
	constructor(p)
	{
		//------------------------------------------------------------
		if(typeof p==='object'){ Object.entries(p).forEach(([key, val]) => { if(typeof val!=='undefined'){ this[key]=val; } }); }
		//------------------------------------------------------------
		if(typeof this.script_name!=='string'){ this.script_name = window.location.href.substring(window.location.href.lastIndexOf('/')+1); }
		//------------------------------------------------------------
		if(typeof this.encoding!=='string'){ this.encoding = null; }
		//------------------------------------------------------------
		if(typeof this.content_type!=='string' || this.content_type===''){ this.content_type = 'text/plain; charset=UTF-8'; }
		//------------------------------------------------------------
		if(typeof this.cache!=='string' || this.cache===''){ this.cache = 'no-store'; }
		//------------------------------------------------------------
		if(typeof this.reject_warnings!=='boolean'){ this.reject_warnings = false; }
		//------------------------------------------------------------
		if(typeof this.debug==='undefined'){ this.debug = false; }
		//------------------------------------------------------------
		this.error = ''; this.warnings = []; this.autoID = 0;
		//------------------------------------------------------------
		this.status = 0; this.status_text = '';
		//------------------------------------------------------------
		this.request = ''; this.response = ''; this.response_headers = null;
		//------------------------------------------------------------
		this.raw_request = ''; this.raw_response='';
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
	// HTTP Functions
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
	send_request({url=null, headers=null, request=null, content_type=null, cache=null, encoding=null, utf8Encode=null}={})
	{
		//------------------------------------------------------------
		this.error = '';
		//------------------------------------------------------------
		if(request===null){ request=''; }
		else
		if(typeof request!=='string'){ request = JSON.stringify(request); }
		//------------------------------------------------------------
		if(typeof content_type!=='string' || content_type===''){ content_type = (typeof this.content_type==='string' && this.content_type!=='') ? this.content_type : 'text/plain; charset=UTF-8'; }
		//------------------------------------------------------------
		if(typeof cache!=='string' || cache===''){ cache = (typeof this.cache==='string' && this.cache!=='') ? this.cache : 'no-store'; }
		//------------------------------------------------------------
		if(typeof encoding!=='string' && typeof this.encoding==='string'){ encoding = this.encoding; }
		//------------------------------------------------------------
		if(typeof utf8Encode!=='boolean'){ utf8Encode = (typeof encoding==='string' && /utf8/i.test(encoding)) ? true : false; }
		//------------------------------------------------------------
		this.status = 0; this.status_text = ''; 
		//------------------------------------------------------------
		this.request = ''; this.response = ''; this.response_headers = null;
		//------------------------------------------------------------
		this.raw_request = request; this.raw_response = '';
		//------------------------------------------------------------
		return new Promise((resolve, reject) =>
		{
			//------------------------------------------------------------
			if(typeof url!=='string' || url==='')
			{
				//------------------------------------------------------------
				return reject('url is not defined or is not a string');
				//------------------------------------------------------------
			}
			else
			{
				//------------------------------------------------------------
				if(utf8Encode===true)
				{
					//----------
					request = this.utf8_encode(request);
					//----------
					if(typeof request!=='string'){ return reject('UTF-8 encoding error'); }
					//---------
				}
				else if(/[^\x00-\xFF]/.test(request)){ return reject('request contains code points greater than 255'); }
				//------------------------------------------------------------
				if(/base64url/i.test(encoding)) // check base64url before base64
				{
					//----------
					request = this.base64url_encode(request, false);
					//----------
					if(typeof request!=='string'){ return reject('base64url encoding error'); }
					//----------
				}
				else
				if(/base64/i.test(encoding)) // check base64 after base64url
				{
					//----------
					request = this.base64_encode(request, false);
					//----------
					if(typeof request!=='string'){ return reject('base64 encoding error'); }
					//----------
				}
				//------------------------------------------------------------
				this.request = request;
				//------------------------------------------------------------
				if(headers===null || typeof headers!=='object'){ headers = {}; }
				//------------------------------------------------------------
				headers['Content-Type'] = content_type;
				//------------------------------------------------------------
				if(typeof headers['Cache-Control']!=='string'){ headers['Cache-Control'] = 'no-store'; }
				if(typeof headers['Pragma']!=='string'){ headers['Pragma'] = 'no-store'; }
				//------------------------------------------------------------
				if(typeof encoding==='string'){ headers['X-Encoding'] = encoding; }
				//------------------------------------------------------------
				if(this.debug===true){ window.console.log(`request: ${request}`); }
				//------------------------------------------------------------
				fetch(url, {method: "POST", redirect: "follow", headers: headers, "body": request, cache: "no-store"})
				//------------------------------------------------------------
				.then(response =>
				{
					//------------------------------------------------------------
					this.status = response.status; this.status_text = response.statusText;
					//------------------------------------------------------------
					this.response_headers = Object.fromEntries(response.headers.entries())
					//------------------------------------------------------------
					if(response.status===200){ return response.text(); }
					else{ return reject('status: '+response.status+': status_text: '+response.statusText); }
					//------------------------------------------------------------
				})
				//------------------------------------------------------------
				.then(data =>
				{
					//------------------------------------------------------------
					if(typeof data==='undefined'){ data = ''; }
					//------------------------------------------------------------
					if(this.debug===true){ window.console.log(`response: ${data}`); }
					//------------------------------------------------------------
					this.raw_response = data;
					//------------------------------------------------------------
					if(encoding==='base64')
					{
						//----------
						data = this.base64_decode(data, false);
						//----------
						if(typeof request!=='string'){ return reject('base64 decoding error'); }
						//----------
					}
					else
					if(encoding==='base64url')
					{
						//----------
						data = this.base64url_decode(data, false);
						//----------
						if(typeof request!=='string'){ return reject('base64url decoding error'); }
						//----------
					}
					//------------------------------------------------------------
					if(utf8Encode)
					{
						//----------
						data = this.utf8_decode(data);
						//----------
						if(typeof data!=='string'){ return reject('UTF-8 decoding error'); }
						//----------
					}
					//------------------------------------------------------------
					this.response = data.trim();
					//------------------------------------------------------------
					return resolve(this.response);
					//------------------------------------------------------------
				})
				//------------------------------------------------------------
				.catch(error => reject(error.message));
				//------------------------------------------------------------
			}
			//------------------------------------------------------------
		});
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
	send_json_request({url=null, request=null, content_type=null, cache=null, encoding=null, utf8Encode=null}={})
	{
		//------------------------------------------------------------
		this.error='';
		//------------------------------------------------------------
		if(typeof content_type!=='string' || content_type===''){ content_type = 'application/json-rpc; charset=UTF-8'; }
		//------------------------------------------------------------
		return new Promise((resolve, reject) =>
		{
			//------------------------------------------------------------
			this.send_request({url, request, content_type, cache, encoding, utf8Encode})
			.then(data =>
			{
				//------------------------------------------------------------
				let json_data = null;
				//------------------------------------------------------------
				try
				{
					//----------
					if(!/^\{.*\}$|^\[.*\]$/.test(data)){ return reject('error parsing json data'); }
					//----------
					json_data = JSON.parse(data);
					//----------
				}
				catch(e1){ return reject('error parsing json data'); }
				//------------------------------------------------------------
				return resolve(json_data);
				//------------------------------------------------------------
			})
			.catch(error => reject(error));
			//------------------------------------------------------------
		});
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
	send_jsonrpc_request({url, request, requests, content_type=null, cache=null, encoding=null, utf8Encode=null}={})
	{
		//------------------------------------------------------------
		this.error=''; this.warnings=[];
		//------------------------------------------------------------
		if(typeof content_type!=='string' || content_type===''){ content_type = 'application/json-rpc; charset=UTF-8'; }
		//------------------------------------------------------------
		return new Promise((resolve, reject) =>
		{
			//------------------------------------------------------------
			if(typeof url!=='string' || url===''){ return reject('url is not defined or is not a string'); }
			if(typeof request==='undefined' && typeof requests==='undefined'){ return reject('request and requests are not defined'); }
			if(typeof request!=='undefined' && ( request instanceof Array || typeof request!=='object' ) ){ return reject('request is not an object'); }
			if(typeof requests!=='undefined' && !(requests instanceof Array)){ return reject('requests is not an array'); }
			//------------------------------------------------------------
			if(request instanceof Object)
			{
				//----------
				if(typeof requests==='undefined'){ requests = []; }
				//----------
				requests.push(request);
				//----------
			}
			//------------------------------------------------------------
			if(requests.length===0){ return reject('requests is an empty array'); }
			//------------------------------------------------------------
			let arrRequests = [];
			//------------------------------------------------------------
			requests.forEach(request =>
			{
				//----------
				if(typeof request!=='object'){ return this.set_warning('request is not defined or is not an object'); }
				//----------
				if(typeof request.method!=='string' || request.method===''){ this.set_warning('method is not defined or is not a string'); }
				if(typeof request.params!=='undefined' && this.datatype(request.params)!=='array' && this.datatype(request.params)!=='object'){ this.set_warning('params is not an array or an object'); }
				if(typeof request.id!=='undefined' && typeof request.id!=='string' && typeof request.id!=='number' && request.id!==null){ this.set_warning('id is not a string, a number or null'); }
				// if(request.id===''){ this.set_warning('id is not defined'); }
				//----------
				let objRequest = {"jsonrpc":"2.0", "method": request.method};
				//----------
				if(typeof request.params!=='undefined'){ objRequest.params = request.params; }
				//----------
				objRequest.id = (typeof request.id!=='undefined') ? request.id : this.auto_id();
				//----------
				arrRequests.push(objRequest);
				//----------
			});
			//------------------------------------------------------------
			if(this.reject_warnings && this.warnings.length===1){ return reject(this.warnings[0]); }
			else if(this.reject_warnings && this.warnings.length>1){ return reject('request contains multiple warnings'); }
			//------------------------------------------------------------
			this.send_request({url, request: arrRequests, content_type, cache, encoding, utf8Encode})
			.then(data =>
			{
				//------------------------------------------------------------
				return resolve(this.parse_jsonrpc_response(data));
				//------------------------------------------------------------
			})
			.catch(error => reject(error));
			//------------------------------------------------------------
		});
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
	// JSON functions
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
	parseJSON(data)
	{
		//----------
		this.error='';
		//----------
		if(typeof data!=='string'){ return this.set_error('data is not defined or is not a string'); }
		//----------
		data = data.trim();
		//----------
		if(data===''){ return null; }
		//----------
		let json_data = null;
		try
		{
			json_data = JSON.parse(data);

		}
		catch(e1)
		{
			return this.set_error('parse error');
		}
		//----------
		return json_data;
		//----------
	}
	//------------------------------------------------------------
	parse_json_response(response)
	{
		//----------
		this.error='';
		//----------
		if(typeof response!=='string'){ return {"error":{"message":'parse error',"data":JSON.stringify(response)}}; }
		//----------
		response = response.trim();
		//----------
		if(response===''){ return {}; }
		//----------
		let json_data = null;
		try
		{
			json_data = JSON.parse(response);

		}
		catch(e1)
		{
			json_data = {"error":{"message":'parse error',"data":JSON.stringify(response)}};
		}
		//----------
		return json_data;
		//----------
	}
	//------------------------------------------------------------
	parse_jsonrpc_response(response)
	{
		//----------
		this.error='';
		//----------
		if(typeof response!=='string'){ return {"jsonrpc":"2.0",
		"error":{"code":-32700,"message":'Parse error',"data":JSON.stringify(response)},
		"id":null}; }
		//----------
		response = response.trim();
		//----------
		if(response===''){ return {}; }
		//----------
		let json_data = null;
		try
		{
			json_data = JSON.parse(response);

		}
		catch(e1)
		{
			json_data = {"jsonrpc":"2.0",
			"error":{"code":-32700,"message":'Parse error',"data":JSON.stringify(response)},
			"id":null};
		}
		//----------
		return json_data;
		//----------
	}
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
	// Encoding / Decoding Functions
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
	base64_encode(data=null, utf8Encode=null)
	{
		//------------------------------------------------------------
		this.error='';
		//------------------------------------------------------------
		if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
		//------------------------------------------------------------
		if(typeof utf8Encode==='boolean' && utf8Encode===true){ data = this.utf8_encode(data); }
		else if(/[^\x00-\xFF]/.test(data)){ return this.set_error('data contains code points greater than 255'); }
		//------------------------------------------------------------
		if(data===''){ return ''; }
		//------------------------------------------------------------
		return btoa(data);
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	base64_decode(data=null, utf8Decode=null)
	{
		//------------------------------------------------------------
		this.error='';
		//------------------------------------------------------------
		if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
		if(/[^A-Za-z0-9+/=]/.test(data)){ return this.set_error('data contains invalid characters'); }
		//------------------------------------------------------------
		if(data===''){ return ''; }
		//------------------------------------------------------------
		try
		{
			data = atob(data);
		}
		catch(e1){ return this.set_error('invalid data'); }
		//----------
		if(typeof utf8Decode==='boolean' && utf8Decode===true){ data = this.utf8_decode(data); }
		//----------
		return data;
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
	base64url_encode(data=null, utf8Encode=null)
	{
		//------------------------------------------------------------
		this.error='';
		//------------------------------------------------------------
		if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
		//------------------------------------------------------------
		if(data===''){ return ''; }
		//------------------------------------------------------------
		data = this.base64_encode(data, utf8Encode);
		//------------------------------------------------------------
		if(data===null){ return null; }
		//------------------------------------------------------------
		return data.replace(/\+/g,'-').replace(/\//g,'_').replace(/\=/g,'');
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	base64url_decode(data=null, utf8Decode=null)
	{
		//------------------------------------------------------------
		this.error='';
		//------------------------------------------------------------
		if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
		//------------------------------------------------------------
		if(data===''){ return ''; }
		//------------------------------------------------------------
		data = data.replace(/\-/g,'+').replace(/\_/g,'/');
		//------------------------------------------------------------
		if(data.length%4){ data += '='.repeat((4-data.length%4)); }
		//------------------------------------------------------------
		return this.base64_decode(data, utf8Decode);
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	//######################################################################
	//------------------------------------------------------------
	base64_base64url(data=null)
	{
		//------------------------------------------------------------
		this.error='';
		//------------------------------------------------------------
		if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
		//------------------------------------------------------------
		if(data===''){ return ''; }
		//------------------------------------------------------------
		return data.replace(/\+/g,'-').replace(/\//g,'_').replace(/\=/g,'');
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	base64url_base64(data=null)
	{
		//------------------------------------------------------------
		this.error='';
		//------------------------------------------------------------
		if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
		//------------------------------------------------------------
		if(data===''){ return ''; }
		//------------------------------------------------------------
		data = data.replace(/\-/g,'+').replace(/\_/g,'/');
		//------------------------------------------------------------
		if(data.length%4){ data += '='.repeat((4-data.length%4)); }
		//------------------------------------------------------------
		return data;
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	//######################################################################
	//------------------------------------------------------------
	utf8_encode(data)
	{
		//------------------------------------------------------------
		this.error='';if(typeof data!=='string'){return this.set_error('data is not defined or is not a string');}if(data===''){return'';}
		//------------------------------------------------------------
		return String.fromCodePoint.apply(null, new TextEncoder().encode(data));
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	utf8_decode(data)
	{
		//------------------------------------------------------------
		this.error='';if(typeof data!=='string'){return this.set_error('data is not defined or is not a string');}if(data===''){return'';}
		//------------------------------------------------------------
		return new TextDecoder().decode(new Uint8Array(Array.from(data).map(s=>s.codePointAt(0))));
		//------------------------------------------------------------
	}
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
	// Other Functions
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
	auto_id(){ if(typeof this.autoID!=='number'){ this.autoID=0; } this.autoID+=1; return this.autoID; }
	//----------
	mt_rand(){ return String(Math.floor(Math.random()*9999999999)).padStart(10,'0'); }
	//----------
	uid1(){ let d1=new Date(); return String(Date.UTC(d1.getFullYear(),d1.getMonth(),d1.getDate(),d1.getHours(),d1.getMinutes(),d1.getSeconds(),d1.getMilliseconds())); }
	//----------
	uid2(){ let d1=new Date(); return (Date.UTC(d1.getFullYear(),d1.getMonth()-1,d1.getDate(),d1.getHours(),d1.getMinutes(),d1.getSeconds(),d1.getMilliseconds())/1000).toFixed(3)+'.'+this.mt_rand(); }
	//------------------------------------------------------------
	datatype(fld)
	{
		//----------
		if(typeof fld==='undefined'){ return 'undefined'; }
		if(fld===null){ return 'null'; }
		if(fld instanceof Array){ return 'array'; }
		if(fld instanceof Date){ return 'date'; }
		//----------
		return typeof fld;
		//----------
	}
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
	// Error Handling Functions
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
	set_warning(warning){ if(typeof warning==='undefined'){ throw new Error('warning is undefined'); } if(typeof this.warnings==='undefined'){ this.warnings=[]; } this.warnings.push(warning); return null; }
	//------------------------------------------------------------
	set_warnings(warnings){ this.warnings = warnings; return null; }
	//------------------------------------------------------------
	get_warnings(){ return this.warnings; }
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
	set_error(error){ if(typeof error==='undefined'){ throw new Error('error is undefined'); } this.error = error; return null; }
	//------------------------------------------------------------
	get_error(){ return this.error; }
	//------------------------------------------------------------
	//############################################################
	//------------------------------------------------------------
}
//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
