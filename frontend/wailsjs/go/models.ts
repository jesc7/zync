export namespace backend {
	
	export class DataPart {
	    key: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new DataPart(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.password = source["password"];
	    }
	}

}

