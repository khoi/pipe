export namespace manifest {
	
	export class Pipe {
	    exec: string;
	    args: string[];
	    stdin?: string;
	
	    static createFrom(source: any = {}) {
	        return new Pipe(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.exec = source["exec"];
	        this.args = source["args"];
	        this.stdin = source["stdin"];
	    }
	}
	export class Manifest {
	    id: string;
	    name: string;
	    description: string;
	    pipe: Pipe;
	    output: string;
	    tags: string[];
	
	    static createFrom(source: any = {}) {
	        return new Manifest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.pipe = this.convertValues(source["pipe"], Pipe);
	        this.output = source["output"];
	        this.tags = source["tags"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

