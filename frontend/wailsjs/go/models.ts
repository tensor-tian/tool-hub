export namespace hub {
	
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

}

export namespace main {
	
	export enum StringValues {
	    ColorOfIcebear = "white",
	    ColorOfMeiMeiBear = "pink",
	    MyFavoriteBear = "MeiMeiBear",
	}
	export enum IntValues {
	    NumberOfBears = 4,
	}

}

