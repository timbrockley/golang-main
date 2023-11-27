/*
################################################################################

Unit Test Class

Copyright (C) 2018-2022	Tim Brockley

The JavaScript code used in this file is free software: you can
redistribute it and/or modify it under the terms of the GNU
General Public License (GNU GPL) as published by the Free Software
Foundation, either version 3 of the License, or (at your option)
any later version.	The code is distributed WITHOUT ANY WARRANTY;
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
export default class UnitTest
//------------------------------------------------------------
{
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    constructor(p)
    {
        //------------------------------------------------------------
        if(typeof p==='object'){ Object.entries(p).forEach(([key, val]) => { if(typeof val!=='undefined'){ this[key]=val; } }); }
        //------------------------------------------------------------
        this.error=''; const that=this;
        //------------------------------------------------------------
        if(typeof this.consoleName!=='string'){ this.consoleName = '_TB_console_'; }
        //------------------------------------------------------------
        this.setup_dom();
        //------------------------------------------------------------
        if(typeof this.sortOrder   !=='string' ){ this.sortOrder    = 'D'; }
        if(typeof this.show_passes !=='boolean'){ this.show_passes  = false; }
        if(typeof this.skip_summary!=='boolean'){ this.skip_summary = false; }
        //------------------------------------------------------------
        this.skipTests = false; this.queueStore = []; this.queueNumber = 0;
        //------------------------------------------------------------
        this.testsPassed = 0; this.testsFailed = 0; this.testsRun = 0;
        //------------------------------------------------------------
        this.time_started = new Date().getTime(); // set here incase but should be replaced by actual tests
        //------------------------------------------------------------
        window.addEventListener('error', function(event)
        {
            //----------
            event.preventDefault();
            //----------
            console.log(event)
            //----------
            that.displayError(event.error, event.lineno, event.filename);
            //----------
        });
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    prepare_test_value(test=null, displayAsString=false)
    {

        //------------------------------------------------------------
        if(test instanceof Error){ if(typeof test==='object' && typeof test.message==='string'){ return `"${test.message}"`; }else{ return `"${test}"`; } }
        //------------------------------------------------------------
        if(typeof test==='string'){ return `"${this.encode_html(test)}"`; }
        //------------------------------------------------------------
        if(this.datatype(test)==='date'){ return `"${this.date_to_string(test)}"`; }
        //------------------------------------------------------------
        const strTest = (this.datatype(test)==='array' || this.datatype(test)==='object') ? JSON.stringify(test) : `${test}`;
        //------------------------------------------------------------
        return (displayAsString===true) ? `"${strTest}"` : strTest;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    prepare_test_results(desc=null, test=null, displayAsString=false)
    {
        //------------------------------------------------------------
        let _results = (desc!==null && desc!=='') ? `${this.encode_html(desc)}: ` : '';
        //------------------------------------------------------------
        const strTest = this.prepare_test_value(test, displayAsString);
        //------------------------------------------------------------
        const style = (test instanceof Error) ? 'color: white; background-color: red; padding: 2px; padding-left: 6px; padding-right: 6px;' : 'color:blue;';
        //------------------------------------------------------------
        return `${_results}[<span style="${style}">${strTest}</span>]`;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    test(desc=null, test=null, displayAsString=false)
    {
        //------------------------------------------------------------
        const _results = this.prepare_test_results(desc, test, displayAsString)
        //------------------------------------------------------------
        let _test = test;
        //------------------------------------------------------------
        if(test===null){ _test = false; }
        else if(typeof test==='string'){ _test = test!==''; }
        else if(typeof test==='number'){ _test = test>0; }
        else if(typeof test==='bigint'){ _test = test>0n; }
        else if(this.datatype(test)==='array'){ _test = test.length>0; }
        else if(this.datatype(test)==='object'){ _test = Object.keys(test).length>0; }
        else{ _test = _test ? true : false; }
        //------------------------------------------------------------
        if(_test===true){ this.pass(_results); }else{ this.fail(_results); }
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    test_string(desc=null, test=null){ return this.test(desc, test, true); }
    //------------------------------------------------------------
    testNE(desc=null, test=null, displayAsString=false)
    {
        //------------------------------------------------------------
        const _results = this.prepare_test_results(desc, test, displayAsString)
        //------------------------------------------------------------
        let _test = test;
        //------------------------------------------------------------
        if(test===null){ _test = true; }
        else if(typeof test==='string'){ _test = test===''; }
        else if(typeof test==='number'){ _test = test<=0; }
        else if(typeof test==='bigint'){ _test = test<=0n; }
        else if(this.datatype(test)==='array'){ _test = test.length===0; }
        else if(this.datatype(test)==='object'){ _test = Object.keys(test).length===0; }
        else{ _test = _test ? false : true; }
        //------------------------------------------------------------
        if(_test===true){ this.pass(_results); }else{ this.fail(_results); }
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    testNE_string(desc=null, test=null){ return this.testNE(desc, test, true); }
    //------------------------------------------------------------
    prepare_compare_results(desc=null, val1=null, val2=null, displayAsString=false)
    {
        //------------------------------------------------------------
        let _results = (desc!==null && desc!=='') ? `${this.encode_html(desc)}: ` : '';
        //------------------------------------------------------------
        const strVal1 = this.prepare_test_value(val1, displayAsString);
        const strVal2 = this.prepare_test_value(val2, displayAsString);
        //------------------------------------------------------------
        const style1 = (val1 instanceof Error) ? 'color: white; background-color: red; padding: 2px; padding-left: 6px; padding-right: 6px;' : 'color:blue;';
        const style2 = (val2 instanceof Error) ? 'color: white; background-color: red; padding: 2px; padding-left: 6px; padding-right: 6px;' : 'color:purple;';
        //------------------------------------------------------------
        return `${_results}[<span style="${style1}">${strVal1}</span>, <span style="${style2}">${strVal2}</span>]`;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    compare(desc=null, val1=null, val2=null, displayAsString=false)
    {
        //------------------------------------------------------------
        const _results = this.prepare_compare_results(desc, val1, val2, displayAsString)
        //------------------------------------------------------------
        let _test = null;
        //------------------------------------------------------------
        if(val1===null || val2===null || typeof val1!=='object' || typeof val2!=='object'){ _test = val1===val2; }
        else
        { _test = this.CompareObject(val1, val2); }
        //------------------------------------------------------------
        if(_test===true){ this.pass(_results); }else{ this.fail(_results); }
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    compare_string(desc=null, val1=null, val2=null){ return this.compare(desc, val1, val2, true); }
    //------------------------------------------------------------
    compareNE(desc=null, val1=null, val2=null, displayAsString=false)
    {
        //------------------------------------------------------------
        const _results = this.prepare_compare_results(desc, val1, val2, displayAsString)
        //------------------------------------------------------------
        let _test = null;
        //------------------------------------------------------------
        if(val1===null || val2===null || typeof val1!=='object' || typeof val2!=='object'){ _test = val1!==val2; }
        else
        { _test = ! this.CompareObject(val1, val2); }
        //------------------------------------------------------------
        if(_test===true){ this.pass(_results); }else{ this.fail(_results); }
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    compareNE_string(desc=null, val1=null, val2=null){ return this.compareNE(desc, val1, val2, true); }
    //------------------------------------------------------------
    compareLE(desc=null, val1=null, val2=null, displayAsString=false)
    {
        //------------------------------------------------------------
        const _results = this.prepare_compare_results(desc, val1, val2, displayAsString)
        //------------------------------------------------------------
        if(val1<=val2){ this.pass(_results); }else{ this.fail(_results); }
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    compareLE_string(desc=null, val1=null, val2=null){ return this.compareLE(desc, val1, val2, true); }
    //------------------------------------------------------------
    compareLT(desc=null, val1=null, val2=null, displayAsString=false)
    {
        //------------------------------------------------------------
        const _results = this.prepare_compare_results(desc, val1, val2, displayAsString)
        //------------------------------------------------------------
        if(val1<val2){ this.pass(_results); }else{ this.fail(_results); }
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    compareLT_string(desc=null, val1=null, val2=null){ return this.compareLT(desc, val1, val2, true); }
    //------------------------------------------------------------
    compareGE(desc=null, val1=null, val2=null, displayAsString=false)
    {
        //------------------------------------------------------------
        const _results = this.prepare_compare_results(desc, val1, val2, displayAsString)
        //------------------------------------------------------------
        if(val1>=val2){ this.pass(_results); }else{ this.fail(_results); }
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    compareGE_string(desc=null, val1=null, val2=null){ return this.compareGE(desc, val1, val2, true); }
    //------------------------------------------------------------
    compareGT(desc=null, val1=null, val2=null, displayAsString=false)
    {
        //------------------------------------------------------------
        const _results = this.prepare_compare_results(desc, val1, val2, displayAsString)
        //------------------------------------------------------------
        if(val1>val2){ this.pass(_results); }else{ this.fail(_results); }
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    compareGT_string(desc=null, val1=null, val2=null){ return this.compareGT(desc, val1, val2, true); }
    //------------------------------------------------------------
    compare_array(desc=null, val1=null, val2=null, displayAsString=false){ return this.compare_object(desc, val1, val2, displayAsString); }
    //------------------------------------------------------------
    compare_array_string(desc=null, val1=null, val2=null){ return this.compare_object(desc, val1, val2, true); }
    //------------------------------------------------------------
    compare_object(desc=null, val1=null, val2=null, displayAsString=false)
    {
        //------------------------------------------------------------
        const _results = this.prepare_compare_results(desc, val1, val2, displayAsString)
        //------------------------------------------------------------
        if(this.CompareObject(val1, val2)){ this.pass(_results); }else{ this.fail(_results); }
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    compare_object_string(desc=null, val1=null, val2=null){ return this.compare_object(desc, val1, val2, true); }
    //------------------------------------------------------------
    pass(str=null)
    {
        //----------------------------------------
        if(typeof str==='undefined'){ str = ''; }
        //----------------------------------------
        const stackTrace = this.stack_trace();
        //----------------------------------------
        this.testsPassed += 1; this.testsRun += 1;
        //----------------------------------------
        str = String(str);
        str = `<div class="pass">${stackTrace.strLineNo}<span style="color:green">PASS</span>: ${str}</div>`;
        //----------------------------------------
        this.writeInnerHTML(str);
        //----------------------------------------
    }
    //------------------------------------------------------------
    fail(str=null)
    {
        //----------------------------------------
        if(typeof str==='undefined'){ str = ''; }
        //----------------------------------------
        const stackTrace = this.stack_trace();
        //----------------------------------------
        this.testsFailed += 1; this.testsRun += 1;
        //----------------------------------------
        str = String(str);
        str = `<div class="fail">${stackTrace.strLineNo}<span style="color:red">FAIL</span>: ${str}</div>`;
        //----------------------------------------
        this.writeInnerHTML(str);
        //----------------------------------------
    }
    //------------------------------------------------------------
    stack_trace()
    {
        //------------------------------------------------------------
        let stack_trace_string='', stack_lines=null, stack_line_matches=null, stack_line_match=null, stack_lineno=null, strLineNo='';
        //------------------------------------------------------------
        stack_trace_string = new Error().stack;
        //------------------------------------------------------------
        if(typeof stack_trace_string==='string')
        {
            if(/stack_trace/.test(stack_trace_string))
            {
                const filename = window.location.href.substring(window.location.href.lastIndexOf('/')+1);
                stack_lines = stack_trace_string.split(/\r\n|\r|\n|\n\r/g);
                if(stack_lines instanceof Array)
                {
                    stack_line_matches=stack_lines.filter(line=>/:\d+:\d+\)*$/gm.test(line) && line.indexOf(filename)>=0)
                    if(stack_line_matches instanceof Array && stack_line_matches.length>0)
                    {
                        stack_line_match = stack_line_matches[0];
                        const line_split = stack_line_match.match(/:(\d+):\d+\)*$/)
                        strLineNo = (line_split instanceof Array && line_split.length===2) ? `(${line_split[1]}) ` : '';
                    }
                }
            }
        }
        else
        {
            stack_trace_string = '';
        }
        //------------------------------------------------------------
        return {stack_trace_string, stack_lines, stack_line_match, stack_lineno, strLineNo};
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    writeInnerHTML(str=null, color=null, backgroundColor=null)
    {
        //----------------------------------------
        if(str===null){ str=''; }
        //----------------------------------------
        if(document.getElementById(this.consoleName))
        {
            //----------
            if(typeof str!=='string'){ str = JSON.stringify(str); }
            //----------
            if(typeof color==='string' || typeof backgroundColor==='string')
            {
                //----------
                color = (typeof color==='string' || color==='') ? color : 'black';
                //----------
                backgroundColor = (typeof backgroundColor==='string' || backgroundColor==='') ? backgroundColor : 'white';
                //----------
                str = `<span style="color:${color};background-color:${backgroundColor}">${str}</span>`;
                //----------
            }
            //----------
            if(typeof color==='string'){ str = `<span style="color:${color}">${str}</span>`; }
            //----------
            if(typeof color==='string'){ str = `<span style="color:${color}">${str}</span>`; }
            //----------
            if(this.sortOrder==='A')
            {
                //----------
                document.getElementById(this.consoleName).innerHTML=document.getElementById(this.consoleName).innerHTML+str;
                //----------
            }
            else
            {
                //----------
                document.getElementById(this.consoleName).innerHTML=str+document.getElementById(this.consoleName).innerHTML;
                //----------
            }
            //----------
        }
        //----------------------------------------
    }
    //----------------------------------------
    write(str=null, color=null, backgroundColor=null)
    {
        //----------------------------------------
        if(str===null){ str=''; }
        //----------------------------------------
        if(typeof str!=='string'){ str=JSON.stringify(str); }
        //----------------------------------------
        str = `<span class="text">${str}</span>`;
        //----------------------------------------
        this.writeInnerHTML(str, color, backgroundColor);
        //----------------------------------------
    }
    //----------------------------------------
    writeln(str=null, color=null, backgroundColor=null)
    {
        //----------------------------------------
        if(str===null){ str=''; }
        //----------------------------------------
        if(typeof str!=='string'){ str=JSON.stringify(str); }
        //----------------------------------------
        str = `<span class="text">${str}</span>`;
        //----------------------------------------
        if(this.sortOrder==='A'){ str='<br>'+str; }else{ str+='<br>'; }
        //----------
        this.write(str, color, backgroundColor);
        //----------------------------------------
    }
    //----------------------------------------
    quoteln(results=null, str=null, color=null, backgroundColor=null)
    {
        //----------------------------------------
        if(results===null){ results = ''; }
        if(str===null){ str=''; }
        //----------------------------------------
        if(typeof results!=='string'){ results = JSON.stringify(results); }
        //----------------------------------------
        results=this.encode_html(results);
        //----------------------------------------
        str=(typeof str==='string') ? JSON.stringify(str) : '"'+JSON.stringify(str)+'"';
        str=this.encode_html(str);
        //----------------------------------------
        this.writeln(`${results}: ${str}`, color, backgroundColor);
        //----------------------------------------
    }
    //------------------------------------------------------------
    quoteln_string(results=null, str=null, color=null, backgroundColor=null)
    {
        //----------------------------------------
        if(results===null){ results = ''; }
        if(str===null){ str=''; }
        //----------------------------------------
        if(typeof results!=='string'){ results = JSON.stringify(results); }
        //----------------------------------------
        results=this.encode_html(results);
        //----------------------------------------
        str=this.encode_html(str);
        //----------------------------------------
        this.writeln(`${results}: "${str}"`, color, backgroundColor);
        //----------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    runTests(tests=null)
    {
        //------------------------------------------------------------
        //############################################################
        //------------------------------------------------------------
        if(tests===null){ return false; }
        //------------------------------------------------------------
        if(!tests instanceof Array){ return false;}
        //------------------------------------------------------------
        //############################################################
        //------------------------------------------------------------
        this.time_started = new Date().getTime();
        //------------------------------------------------------------
        let testNumber=0;
        //------------------------------------------------------------
        //############################################################
        //------------------------------------------------------------
        while(testNumber<tests.length)
        {
            //----------
            if(this.skipTests===true){ break; }
            //----------
            testNumber+=1;
            //----------
            if(typeof tests[testNumber-1]!=='undefined')
            {
                //----------
                this.write(`<hr>`);
                //----------
                const timeStart = new Date().getTime();
                //----------
                tests[testNumber-1]();
                //----------
                const timeDif = new Date().getTime() - timeStart;
                //----------
                this.write(`<hr>ms = ${timeDif}<hr>`);
                //----------
            }
            //----------
        }
        //------------------------------------------------------------
        //############################################################
        //------------------------------------------------------------
        if(!this.skip_summary){ this.displayTestResults(); }
        //------------------------------------------------------------
        //############################################################
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    displayError(error=null, lineno=null, filename=null)
    {
        //----------------------------------------
        console.error(error);
        //----------------------------------------
        if(typeof error==='undefined' || error===null){ error='undefined error'; }
        //----------------------------------------
        let strErrorInfo = '';
        if(typeof lineno==='number'){ strErrorInfo += `(${lineno}) `; }
        if(typeof filename==='string'){ strErrorInfo += filename; }
        //----------------------------------------
        const highlightColor = (lineno) ? 'red' : 'orange';
        //----------------------------------------
        let strError = `<div class="fail" style="margin: 10px; margin-left: 0px; margin-right: 0px;"><span style="background-color: ${highlightColor};color: white; font-weight: bold; padding: 4px; padding-left: 6px; padding-right: 6px;">${error}</span>`;
        if(strErrorInfo!==''){ strError += `<div style="color: ${highlightColor}; font-weight: bold;padding-top: 12px;">${strErrorInfo}</div></div>`; }
        strError += `</div>`;
        //----------------------------------------
        this.writeInnerHTML(strError);
        //----------------------------------------
    }
    //----------------------------------------
    //########################################
    //----------------------------------------
    displayTestResults()
    {
        //----------------------------------------
        let str=`<br><a href="" onclick="togglePass('pass'); return false;">PASSED = ${this.testsPassed}</a>`;
        //----------------------------------------
        if(this.testsFailed)
        {
            str+=` <a href="" onclick="toggleFail('fail'); return false;">FAILED = ${this.testsFailed}</a>`;
            str+=` total tests = ${this.testsRun}`;
        }
        //----------------------------------------
        this.time_ended = new Date().getTime(); this.time_taken = this.time_ended - this.time_started;
        //----------------------------------------
        str+=' ('+this.time_taken+' ms)';
        //----------------------------------------
        str+='<br><br>'; this.writeInnerHTML(str);
        //----------------------------------------
        if(this.show_passes===false && this.testsFailed>0){ showClassOnly('fail'); }
        //----------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    skip_tests(display_test_results)
    {
        this.skipTests=true;
        if(typeof display_test_results==='boolean' && display_test_results===true){ this.displayTestResults(); }
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    // queing system (ie. run tasks one at a time and pause until requested - see TB.queue)
    //------------------------------------------------------------
    queue(queue=null, ms=null)
    {
        //------------------------------------------------------------
        if(queue===null){ throw 'queue not defined'; }
        //------------------------------------------------------------
        this.time_started = new Date().getTime();
        //------------------------------------------------------------
        if(typeof queue!=='function' && !queue instanceof Array){ throw 'queue not a function or and an instance of an array';}
        //------------------------------------------------------------
        try{ this.queueStore=this.queueStore.concat(queue); }catch(e1){}
        //------------------------------------------------------------
        if(typeof ms==='number'){ this.run_next(ms); }
        //------------------------------------------------------------
        return this;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    run_next(ms=null)
    {
        //------------------------------------------------------------
        const that = this;
        //------------------------------------------------------------
        if(this.queueNumber>=this.queueStore.length){ throw 'end of queue reached'; }
        //------------------------------------------------------------
        if(this.skipTests===true){ return null; }
        //------------------------------------------------------------
        this.queueNumber+=1;
        //------------------------------------------------------------
        if(typeof ms==='number')
        {
            window.setTimeout(function(){ that.queueStore[that.queueNumber-1].call(that); }, ms);
        }
        else
        {
            this.queueStore[this.queueNumber-1].call(that);
        }
        //------------------------------------------------------------
        return that;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    encodeHTML(str=null)
    {
        //------------------------------------------------------------
        if(typeof str!=='string' || str===''){ return ''; }
        //------------------------------------------------------------
        str=str.replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/"/g,'&quot;').replace(/ /g,'&nbsp;');
        str=str.replace(/\r\n|\n\r|\r|\n/g,'<br>');
        //------------------------------------------------------------
        return str;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    encode_html(str=null)
    {
        //------------------------------------------------------------
        if(typeof str!=='string' || str===''){ return ''; }
        //------------------------------------------------------------
        str=str.replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/"/g,'&quot;').replace(/ /g,'&nbsp;');
        //------------------------------------------------------------
        return str;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    CompareObject(obj1=null, obj2=null){ return (JSON.stringify(obj1)===JSON.stringify(obj2)); }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    setup_dom()
    {
        //------------------------------------------------------------
        window.top.document.title = window.location.href.substring(window.location.href.lastIndexOf('/')+1);
        //------------------------------------------------------------
        if(typeof this.consoleName!=='string'){ this.consoleName='_TB_console_'; }
        //------------------------------------------------------------
        const styleSheet=`#${this.consoleName} { font-family: monospace; font-size: larger; } .hidden { display: none; } .classBackground1 { background-color: red; } .classBackground2 { background-color: blue; }`;
        //------------------------------------------------------------
        const style = document.createElement('style');
        style.innerHTML = styleSheet;
        document.head.appendChild(style);
        //------------------------------------------------------------
        if(document.getElementById(this.consoleName)===null)
        {
            let divConsole=document.createElement('div');
            divConsole.setAttribute('id', this.consoleName);
            document.body.prepend(divConsole);
        }
        //------------------------------------------------------------
        let html_scripts=
        `
        //------------------------------------------------------------
        let showingClassOnly = null;
        //------------------------------------------------------------
        function togglePass()
        {
            //----------
            if(showingClassOnly==='pass'){ showAllClasses(); }
            else{ showClassOnly('pass'); }
            //----------
            return false;
            //----------
        }
        //------------------------------------------------------------
        function toggleFail()
        {
            //----------
            if(showingClassOnly==='fail'){ showAllClasses(); }
            else{ showClassOnly('fail'); }
            //----------
            return false;
            //----------
        }
        //-----------------------------------------------------------
        function showAllClasses()
        {
            //----------
            showingClassOnly = null;
            //----------
            document.querySelectorAll('.pass, .fail, .text').forEach(element => { element.classList.remove('hidden'); });
            //----------
            return false;
            //----------
        }
        //------------------------------------------------------------
        function showClassOnly(className)
        {
            //----------
            showingClassOnly = className;
            //----------
            document.querySelectorAll('.pass, .fail, .text')
            .forEach(element =>
            {
                if(element.classList.contains(className))
                { element.classList.remove('hidden'); }
                else{ element.classList.add('hidden'); }
            });
            //----------
            return false;
            //----------
        }
        //------------------------------------------------------------
        `;
        //----------------------------------------
        let s = document.createElement('script');
        let inlineScript = document.createTextNode(html_scripts);
        s.appendChild(inlineScript);
        document.head.appendChild(s);
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    datatype(fld)
    {
        //------------------------------------------------------------
        if(typeof fld==='undefined'){ return 'undefined'; }
        if(fld===null){ return 'null'; }
        if(fld instanceof Array){ return 'array'; }
        if(fld instanceof Date){ return 'date'; }
        //------------------------------------------------------------
        return typeof fld;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    date_to_string(d1)
    {
        //----------
        if(typeof d1==='undefined'){ d1=new Date(); }
        //----------
        return d1.getUTCFullYear()+'-'+String(d1.getUTCMonth()+1).padStart(2,'0')+'-'+String(d1.getUTCDate()).padStart(2,'0')+'T'+String(d1.getUTCHours()).padStart(2,'0')+':'+String(d1.getUTCMinutes()).padStart(2,'0')+':'+String(d1.getUTCSeconds()).padStart(2,'0')+'.'+String(d1.getUTCMilliseconds()).padStart(3,'0')+'Z';
        //----------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    trim(data){ if(typeof data!=='string'||data===''){ return ''; } return data.replace(/^\s*/,'').replace(/\s*$/,''); }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    stringToLongs(data=null)
    {
        //-------------------------------------------------------------
        // converts a string of 8 bit characters (in 4 characters chunks) to 32 bit longs
        //-------------------------------------------------------------
        if(typeof data!=='string'){ return this.set_error('data is not defined or is not a string'); }
        if(/[^\x00-\xFF]/.test(data)){ return this.set_error('data contains code points greater than 255'); }
        //-------------------------------------------------------------
        if(data===''){ return [0]; }
        //-------------------------------------------------------------
        // should be a multiple of 4 bytes
        if(data.length % 4){ data=data.padEnd(data.length+(4 - (data.length % 4)), "\0"); }
        //----------
        let longs = new Array(Math.ceil(data.length/4));
        //----------
        for(let i=0; i<longs.length; i+=1)
        {
            longs[i] = data.charCodeAt(i*4)
            | (( data.charCodeAt(i*4+1) << 8  )>>>0)
            | (( data.charCodeAt(i*4+2) << 16 )>>>0)
            | (( data.charCodeAt(i*4+3) << 24 )>>>0);
        }
        //----------
        return longs;
        //----------
    }
    //------------------------------------------------------------
    longsToString(longs=null, rightTrimNull=null)
    {
        //----------
        if(!(longs instanceof Array)){ return this.set_error('longs is not defined or is not an array'); }
        //----------
        if(typeof rightTrimNull!=='boolean'){ rightTrimNull = false; }
        //----------
        let data = '';
        //----------
        for(let i=0; i<longs.length; i++)
        {
            //----------
            if(!Number.isInteger(longs[i])){ return this.set_error('longs cotains non integers'); }
            //----------
            data += String.fromCharCode(longs[i] & 0xFF, longs[i] >>> 8 & 0xFF, longs[i] >>> 16 & 0xFF, longs[i] >>> 24 & 0xFF);
            //----------
        }
        //----------
        return rightTrimNull ? data.replace(/\0*$/g, '') : data;
        //----------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    ascii_encode({data=null, utf8Encode=null, escapeString=null}={})
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
        //------------------------------------------------------------
        if(typeof utf8Encode==='boolean' && utf8Encode===true){ data = this.utf8_encode(data); }
        //------------------------------------------------------------
        data = data.replace(/\\/g, '\\\\');
        //------------------------------------------------------------
        if(typeof escapeString==='boolean' && escapeString===true)
        {
            data = data.replace(/[\x22\x27\x60]/g,
            function(char)
            {
                return '\\x'+char.codePointAt(0).toString(16).toUpperCase().padStart(2, '0');
            });
        }
        //------------------------------------------------------------
        return data
        .replace(/[\x00-\x1F\x7F-\xFF]/g,
        function(char)
        {
            return '\\x'+char.codePointAt(0).toString(16).toUpperCase().padStart(2, '0');
        })
        .replace(/[\u{0100}-\u{10FFFF}]/gu,
        function(char)
        {
            return '\\u{'+char.codePointAt(0).toString(16).padStart(4, '0')+'}';
        });
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    ascii_decode({data=null, utf8Decode=null}={})
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
        //------------------------------------------------------------
        if(data===''){ return ''; }
        //------------------------------------------------------------
        data = data.replace(/\\\\/g, '\\SUB')
        .replace(/\\x[0-9A-F]{2}|\\u\{[0-9A-F]{4,6}\}/gi,
        function(hex)
        {
            return String.fromCodePoint(parseInt(hex.replace(/[\\x\\u{}]/gi,''), 16));

        })
        .replace(/\\SUB/g, '\\');
        //------------------------------------------------------------
        if(typeof utf8Decode==='boolean' && utf8Decode===true){ data = this.utf8_decode(data); }
        //------------------------------------------------------------
        return data;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    base32_encode(data=null, utf8Encode=null)
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(typeof data!=='string'){ return this.set_error('data is not defined or is not a string'); }
        //------------------------------------------------------------
        if(typeof utf8Encode==='boolean' && utf8Encode===true){ data = this.utf8_encode(data); }
        else if(/[^\x00-\xFF]/.test(data)){ return this.set_error('data contains code points greater than 255'); }
        //------------------------------------------------------------
        if(data===''){ return ''; }
        //------------------------------------------------------------
        let binString = '', encString = '';
        //------------------------------------------------------------
        for(let char of data)
        {
            //----------
            binString += char.codePointAt(0).toString(2).padStart(8,'0');
            //----------
        };
        //------------------------------------------------------------
        // pad binary string with zeros if last 5 bit chunk length was less than 5
        if(binString.length % 5){ binString = binString.padEnd(binString.length+(5 - (binString.length % 5)), "0"); }
        //------------------------------------------------------------
        // map 5 bit chunks to character encodings
        binString.match(/.{1,5}/g).forEach(chunk =>
        {
            //----------
            encString += this.base32_map[parseInt(chunk, 2)];
            //----------
        });
        //------------------------------------------------------------
        // pad encoded string if not a multiple of 8
        if(encString.length % 8){ encString = encString.padEnd(encString.length+(8 - (encString.length % 8)), "="); }
        //------------------------------------------------------------
        return encString;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    base32_decode(data=null, utf8Decode=null)
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(typeof data!=='string'){ return this.set_error('data is not defined or is not a string'); }
        if(/[^A-Z2-7=]/.test(data)){ return this.set_error('data contains invalid characters'); }
        //------------------------------------------------------------
        data = data.replace(/=/g, '');
        //------------------------------------------------------------
        if(data===''){ return ''; }
        //------------------------------------------------------------
        let binString = '', decString = '';
        //------------------------------------------------------------
        // decode character string into 5 bit chunks
        for(let char of data)
        {
            //----------
            const mapIndex = this.base32_map.indexOf(char);
            const mappedCodePoint = (mapIndex >= 0) ? mapIndex : 0;
            //----------
            binString += mappedCodePoint.toString(2).padStart(5,'0');
            //----------
        }
        //------------------------------------------------------------
        // convert code points to character string (ignore zero padding)
        const strLen = Math.floor(binString.length/8)*8;
        let i = 0;
        while(i < strLen)
        {
            //----------
            decString += String.fromCodePoint(parseInt(binString.substring(i,i+8), 2));
            //----------
            i+=8;
            //----------
        }
        //------------------------------------------------------------
        if(typeof utf8Decode==='boolean' && utf8Decode===true){ decString = this.utf8_decode(decString); }
        //------------------------------------------------------------
        return decString;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //##################################################
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
        //------------------------------------------------------------
        if(typeof utf8Decode==='boolean' && utf8Decode===true){ data = this.utf8_decode(data); }
        //------------------------------------------------------------
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
    bytesToBase64(bytesArray=null)
    {
        //------------------------------------------------------------
        this.error = '';
        //------------------------------------------------------------
        if(!(bytesArray instanceof Array)){ return this.set_error('bytesArray is not defined or is not an array'); }
        //------------------------------------------------------------
        return this.base64_encode(this.codepoints_decode(bytesArray));
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    bytesToBase64url(bytesArray=null)
    {
        //------------------------------------------------------------
        return this.base64_base64url(this.bytesToBase64(bytesArray));
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    base64ToBytes(base64Str=null)
    {
        //------------------------------------------------------------
        this.error = '';
        //------------------------------------------------------------
        if(typeof base64Str!=='string'){ return this.set_error('base64Str is not defined or is not a string'); }
        if(/[^A-Za-z0-9+/=]/.test(base64Str)){ return this.set_error('base64Str contains invalid characters'); }
        //------------------------------------------------------------
        return this.codepoints_encode(this.base64_decode(base64Str));
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    base64urlToBytes(base64Str=null)
    {
        //------------------------------------------------------------
        return this.base64ToBytes(this.base64url_base64(base64Str));
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    hex_encode(data=null, utf8Encode=null)
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
        let strArray = [];
        //------------------------------------------------------------
        for(let i = 0; i<data.length; i++)
        {
            strArray.push(data.charCodeAt(i).toString(16).padStart(2,'0'));
        }
        //------------------------------------------------------------
        return strArray.join('').toUpperCase();
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    hex_decode(data=null, utf8Decode=null)
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
        if(data.length % 2){ return this.set_error('data length is not a multiple of 2'); }
        if(/[^0-9A-F]/i.test(data)){ return this.set_error('data contains invalid characters'); }
        //------------------------------------------------------------
        if(data===''){ return ''; }
        //------------------------------------------------------------
        let strArray = [];
        //------------------------------------------------------------
        for(let i=0; i<data.length-1; i+=2)
        {
            strArray.push(parseInt(data.substring(i, i+2), 16));
        }
        //------------------------------------------------------------
        data = String.fromCharCode.apply(String, strArray);
        //------------------------------------------------------------
        if(typeof utf8Decode==='boolean' && utf8Decode===true){ data = this.utf8_decode(data); }
        //------------------------------------------------------------
        return data;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    octal_encode(data=null, utf8Encode=null)
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
        let strArray = [];
        //------------------------------------------------------------
        for(let i = 0; i<data.length; i++)
        {
            strArray.push(data.charCodeAt(i).toString(8).padStart(3,'0'));
        }
        //------------------------------------------------------------
        return strArray.join('');
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    octal_decode(data=null, utf8Decode=null)
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
        if(data.length % 3){ return this.set_error('data length is not a multiple of 3'); }
        if(/[^0-7]/i.test(data)){ return this.set_error('data contains invalid characters'); }
        //------------------------------------------------------------
        if(data===''){ return ''; }
        //------------------------------------------------------------
        let strArray = [];
        //------------------------------------------------------------
        for(let i=0; i<data.length-1; i+=3)
        {
            strArray.push(parseInt(data.substring(i, i+3), 8));
        }
        //------------------------------------------------------------
        data = String.fromCharCode.apply(String, strArray);
        //------------------------------------------------------------
        if(typeof utf8Decode==='boolean' && utf8Decode===true){ data = this.utf8_decode(data); }
        //------------------------------------------------------------
        return data;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    charcodes_encode(data, utf8Encode)
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
        //------------------------------------------------------------
        if(typeof utf8Encode==='boolean' && utf8Encode===true){ data = this.utf8_encode(data); }
        //------------------------------------------------------------
        if(data===''){ return []; }
        //------------------------------------------------------------
        let charCodes = [];
        //----------
        for(let i=0; i<data.length; i+=1)
        {
            charCodes.push(data[i].charCodeAt(0));
        }
        //------------------------------------------------------------
        return charCodes;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    charcodes_decode(charCodes, utf8Decode)
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(this.datatype(charCodes) !== 'array'){ return this.set_error('charCodes is not defined or is not an array'); }
        //------------------------------------------------------------
        if(charCodes.length===0){ return ''; }
        //------------------------------------------------------------
        let data = '';
        //----------
        for(let i=0; i<charCodes.length; i++)
        {
            data += String.fromCharCode(charCodes[i])
        }
        //------------------------------------------------------------
        if(typeof utf8Decode==='boolean' && utf8Decode===true){ data = this.utf8_decode(data); }
        //------------------------------------------------------------
        return data;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    codepoints_encode(data, utf8Encode)
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
        //------------------------------------------------------------
        if(typeof utf8Encode==='boolean' && utf8Encode===true){ data = this.utf8_encode(data); }
        //------------------------------------------------------------
        if(data===''){ return []; }
        //------------------------------------------------------------
        let codePoints = [];
        //----------
        for(let char of data)
        {
            codePoints.push(char.codePointAt(0));
        }
        //------------------------------------------------------------
        return codePoints;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    codepoints_decode(codePoints, utf8Decode)
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(this.datatype(codePoints) !== 'array'){ return this.set_error('codePoints is not defined or is not an array'); }
        //------------------------------------------------------------
        if(codePoints.length===0){ return ''; }
        //------------------------------------------------------------
        let data = String.fromCodePoint.apply(null, codePoints);
        //------------------------------------------------------------
        if(typeof utf8Decode==='boolean' && utf8Decode===true){ data = this.utf8_decode(data); }
        //------------------------------------------------------------
        return data;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    escape_stringV1(data)
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
        //------------------------------------------------------------
        if(data===''){ return ''; }
        //------------------------------------------------------------
        return data.replace(/\\/g,'\\\\').replace(/\x09/g,'\\t').replace(/\x0A/g,'\\n').replace(/\x0D/g,'\\r').replace(/\x22/g,'\\q').replace(/\x27/g,'\\a').replace(/\x60/g,'\\g');
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    unescape_stringV1(data)
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
        //------------------------------------------------------------
        if(data===''){ return ''; }
        //------------------------------------------------------------
        return data.replace(/\\\\/g,'\\SUB').replace(/\\g/g,'\x60').replace(/\\a/g,'\x27').replace(/\\q/g,'\x22').replace(/\\r/g,'\x0D').replace(/\\n/g,'\x0A').replace(/\\t/g,'\x09').replace(/\\SUB/g,'\\');
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    escape_stringV2(data)
    {
        //------------------------------------------------------------
        if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
        //------------------------------------------------------------
        if(data===''){ return ''; }
        //------------------------------------------------------------
        return data.replace(/~/g,'~~').replace(/\x09/g,'~t').replace(/\x0A/g,'~n').replace(/\x0D/g,'~r').replace(/\x22/g,'~q').replace(/\x27/g,'~a').replace(/\x5C/g,'~b').replace(/\x60/g,'~g');
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    unescape_stringV2(data)
    {
        //------------------------------------------------------------
        if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
        //------------------------------------------------------------
        if(data===''){ return ''; }
        //------------------------------------------------------------
        return data.replace(/~~/g,'~SUB').replace(/~g/g,'\x60').replace(/~b/g,'\x5C').replace(/~a/g,'\x27').replace(/~q/g,'\x22').replace(/~r/g,'\x0D').replace(/~n/g,'\x0A').replace(/~t/g,'\x09').replace(/~SUB/g,'~');
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    text_encodeV1(data=null, utf8Encode=null)
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
        return data
        .replace(/[\x00-\x1F\x22\x25\x26\x27]/g,
        function(char)
        {
            return '\x25'+String.fromCodePoint(char.codePointAt(0)+40);
        })
        .replace(/[\x7F-\xAD]/g,
        function(char)
        {
            return '\x25'+String.fromCodePoint(char.codePointAt(0)-47);
        })
        .replace(/[\xAE-\xFF]/g,
        function(char)
        {
            return '\x26'+String.fromCodePoint(char.codePointAt(0)-134);
        })
        .replace(/\x5C/g, '\x26\x7A')
        .replace(/\x60/g, '\x26\x7B');
            //------------------------------------------------------------
    }
    //------------------------------------------------------------
    text_decodeV1(data=null, utf8Decode=null)
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
        //------------------------------------------------------------
        if(data===''){ return ''; }
        //------------------------------------------------------------
        data = data
        .replace(/\x26\x7B/g, '\x60')
        .replace(/\x26\x7A/g, '\x5C')
            .replace(/\x26[\x28-\x7E]/g,
        function(char)
        {
            return String.fromCodePoint(char.codePointAt(1)+134);

        })
        .replace(/\x25[\x50-\x7E]/g,
        function(char)
        {
            return String.fromCodePoint(char.codePointAt(1)+47);

        })
        .replace(/\x25[\x28-\x4F]/g,
        function(char)
        {
            return String.fromCodePoint(char.codePointAt(1)-40);

        });
        //------------------------------------------------------------
        if(typeof utf8Decode==='boolean' && utf8Decode===true){ data = this.utf8_decode(data); }
        //------------------------------------------------------------
        return data;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
    unicode_encode({data=null, utf8Encode=null, escapeString=null}={})
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
        //------------------------------------------------------------
        if(typeof utf8Encode==='boolean' && utf8Encode===true){ data = this.utf8_encode(data); }
        //------------------------------------------------------------
        if(data===''){ return ''; }
        //------------------------------------------------------------
        data = data.replace(/\\/g, '\\\\');
        //------------------------------------------------------------
        if(typeof escapeString==='boolean' && escapeString===true)
        {
            data = data.replace(/[\x22\x27\x60]/g,
            function(char)
            {
                return '\\u{'+char.codePointAt(0).toString(16).padStart(4, '0')+'}';
            });
        }
        //------------------------------------------------------------
        return data
        .replace(/[\x00-\x1F\x7F-\xFF\u{0100}-\u{10FFFF}]/gu,
        function(char)
        {
            return '\\u{'+char.codePointAt(0).toString(16).padStart(4, '0')+'}';
        });
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    unicode_decode({data=null, utf8Decode=null}={})
    {
        //------------------------------------------------------------
        this.error='';
        //------------------------------------------------------------
        if(typeof data !== 'string'){ return this.set_error('data is not defined or is not a string'); }
        //------------------------------------------------------------
        if(data===''){ return ''; }
        //------------------------------------------------------------
        data = data.replace(/\\\\/g, '\\SUB')
        .replace(/\\u\{[0-9A-F]{4,6}\}/gi,
        function(hex)
        {
            return String.fromCodePoint(parseInt(hex.replace(/[\\u{}]/gi,''), 16));

        })
        .replace(/\\SUB/g, '\\');
        //------------------------------------------------------------
        if(typeof utf8Decode==='boolean' && utf8Decode===true){ data = this.utf8_decode(data); }
        //------------------------------------------------------------
        return data;
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
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
    set_error(error)
    {
        //------------------------------------------------------------
        if(typeof error!=='string'){ error='undefined error'; }
        //------------------------------------------------------------
        return this.displayError(error);
        //------------------------------------------------------------
    }
    //------------------------------------------------------------
    //############################################################
    //------------------------------------------------------------
}
//------------------------------------------------------------
//############################################################
//------------------------------------------------------------
//############################################################
