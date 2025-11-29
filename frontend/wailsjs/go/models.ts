export namespace hub {
	
	export enum StringValues {
	    SettingKeyToolsDir = "ToolsDir",
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
	    name: string;
	    description: string;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new Tool(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	        this.type = source["type"];
	    }
	}
	export class RespGetToolList {
	    error: string;
	    list: Tool[];
	
	    static createFrom(source: any = {}) {
	        return new RespGetToolList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	        this.list = this.convertValues(source["list"], Tool);
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
	export class ToolDetail {
	    name: string;
	    description: string;
	    type: string;
	    parameters: string;
	    logLifeSpan: string;
	    concurrencyGroupName: string;
	    extra: string;
	
	    static createFrom(source: any = {}) {
	        return new ToolDetail(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	        this.type = source["type"];
	        this.parameters = source["parameters"];
	        this.logLifeSpan = source["logLifeSpan"];
	        this.concurrencyGroupName = source["concurrencyGroupName"];
	        this.extra = source["extra"];
	    }
	}
	export class RespToolDetail {
	    error: string;
	    // Go type: ToolDetail
	    item: any;
	
	    static createFrom(source: any = {}) {
	        return new RespToolDetail(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	        this.item = this.convertValues(source["item"], null);
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
	

}

export namespace main {
	
	export enum IntValues {
	    NumberOfBears = 4,
	}

}

