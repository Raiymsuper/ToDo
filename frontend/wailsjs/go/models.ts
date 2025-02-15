export namespace main {
	
	export class todos {
	    Id: number;
	    Task: string;
	    Status: boolean;
	
	    static createFrom(source: any = {}) {
	        return new todos(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Task = source["Task"];
	        this.Status = source["Status"];
	    }
	}

}

