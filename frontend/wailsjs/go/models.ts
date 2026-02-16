export namespace backend {
	
	export class OfferData {
	    value: string;
	    key: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new OfferData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.value = source["value"];
	        this.key = source["key"];
	        this.password = source["password"];
	    }
	}

}

