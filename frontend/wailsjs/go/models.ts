export namespace hub {
	
	export enum StringValues {
	    SettingKeyToolsDir = "ToolsDir",
	}
	export class CommandLineToolExtra {
	    sh: string;
	    wd: string;
	    cmd: string;
	    env: string;
	    stdin: string;
	
	    static createFrom(source: any = {}) {
	        return new CommandLineToolExtra(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.sh = source["sh"];
	        this.wd = source["wd"];
	        this.cmd = source["cmd"];
	        this.env = source["env"];
	        this.stdin = source["stdin"];
	    }
	}
	export class CommandLineTool {
	    id: number;
	    createdAt: number;
	    updatedAt: number;
	    name: string;
	    description: string;
	    parameters: string;
	    category: string;
	    schema: string;
	    definition: string;
	    code: string;
	    defaultParams: string;
	    logLifeSpan: string;
	    concurrencyGroupName: string;
	    timeout: string;
	    isStream: boolean;
	    extra: CommandLineToolExtra;
	
	    static createFrom(source: any = {}) {
	        return new CommandLineTool(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.parameters = source["parameters"];
	        this.category = source["category"];
	        this.schema = source["schema"];
	        this.definition = source["definition"];
	        this.code = source["code"];
	        this.defaultParams = source["defaultParams"];
	        this.logLifeSpan = source["logLifeSpan"];
	        this.concurrencyGroupName = source["concurrencyGroupName"];
	        this.timeout = source["timeout"];
	        this.isStream = source["isStream"];
	        this.extra = this.convertValues(source["extra"], CommandLineToolExtra);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class Dirs {
	    home: string;
	    temp: string;
	    app: string;
	
	    static createFrom(source: any = {}) {
	        return new Dirs(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.home = source["home"];
	        this.temp = source["temp"];
	        this.app = source["app"];
	    }
	}
	export class HTTPToolExtra {
	    url: string;
	    method: string;
	    query: string;
	    headers: Record<string, string>;
	    body: string;
	
	    static createFrom(source: any = {}) {
	        return new HTTPToolExtra(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.method = source["method"];
	        this.query = source["query"];
	        this.headers = source["headers"];
	        this.body = source["body"];
	    }
	}
	export class HTTPTool {
	    id: number;
	    createdAt: number;
	    updatedAt: number;
	    name: string;
	    description: string;
	    parameters: string;
	    category: string;
	    schema: string;
	    definition: string;
	    code: string;
	    defaultParams: string;
	    logLifeSpan: string;
	    concurrencyGroupName: string;
	    timeout: string;
	    isStream: boolean;
	    extra: HTTPToolExtra;
	
	    static createFrom(source: any = {}) {
	        return new HTTPTool(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.parameters = source["parameters"];
	        this.category = source["category"];
	        this.schema = source["schema"];
	        this.definition = source["definition"];
	        this.code = source["code"];
	        this.defaultParams = source["defaultParams"];
	        this.logLifeSpan = source["logLifeSpan"];
	        this.concurrencyGroupName = source["concurrencyGroupName"];
	        this.timeout = source["timeout"];
	        this.isStream = source["isStream"];
	        this.extra = this.convertValues(source["extra"], HTTPToolExtra);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class RespGetCommandLineTool {
	    error: string;
	    item: CommandLineTool;
	
	    static createFrom(source: any = {}) {
	        return new RespGetCommandLineTool(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	        this.item = this.convertValues(source["item"], CommandLineTool);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RespGetHTTPTool {
	    error: string;
	    item: HTTPTool;
	
	    static createFrom(source: any = {}) {
	        return new RespGetHTTPTool(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	        this.item = this.convertValues(source["item"], HTTPTool);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RespGetSettings {
	    error: string;
	    kvMap: Record<string, string>;
	
	    static createFrom(source: any = {}) {
	        return new RespGetSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	        this.kvMap = source["kvMap"];
	    }
	}
	export class Tool {
	    id: number;
	    createdAt: number;
	    updatedAt: number;
	    name: string;
	    description: string;
	    parameters: string;
	    category: string;
	    schema: string;
	    definition: string;
	    code: string;
	    defaultParams: string;
	
	    static createFrom(source: any = {}) {
	        return new Tool(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.parameters = source["parameters"];
	        this.category = source["category"];
	        this.schema = source["schema"];
	        this.definition = source["definition"];
	        this.code = source["code"];
	        this.defaultParams = source["defaultParams"];
	    }
	}
	export class RespGetTool {
	    error: string;
	    item: Tool;
	
	    static createFrom(source: any = {}) {
	        return new RespGetTool(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	        this.item = this.convertValues(source["item"], Tool);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ToolBrief {
	    id: number;
	    name: string;
	    description: string;
	    category: string;
	
	    static createFrom(source: any = {}) {
	        return new ToolBrief(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.category = source["category"];
	    }
	}
	export class RespGetToolList {
	    error: string;
	    list: ToolBrief[];
	
	    static createFrom(source: any = {}) {
	        return new RespGetToolList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	        this.list = this.convertValues(source["list"], ToolBrief);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ToolTestcase {
	    id: number;
	    createdAt: number;
	    updatedAt: number;
	    toolName: string;
	    input: string;
	    output: string;
	    ok: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ToolTestcase(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.toolName = source["toolName"];
	        this.input = source["input"];
	        this.output = source["output"];
	        this.ok = source["ok"];
	    }
	}
	export class RespGetToolTestcaseList {
	    error: string;
	    list: ToolTestcase[];
	
	    static createFrom(source: any = {}) {
	        return new RespGetToolTestcaseList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	        this.list = this.convertValues(source["list"], ToolTestcase);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RespSaveSetting {
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new RespSaveSetting(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	    }
	}
	
	

}

export namespace main {
	
	export enum IntValues {
	    NumberOfBears = 4,
	}

}

