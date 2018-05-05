from jsonschema import validate

schema = {
	"colors" : {
		"color" : {"type" : "string"},
		"category" : {"type" : "string"},
		"type" : {"type" : "string"},
		"code" : {
			"rgba" : {"type" : "array"},
			"hex" : {"type" : "string"},
			"required" : ["rgba", "hex"]					
		},
		"required" : ["color", "category", "type", "code"]
	},
	"required" : ["colors"]
}

# schema = {
# 	"colors" : {
# 		"type" : "object",
# 		"properties" : {
# 			"color" : {"type" : "string"},
# 			"category" : {"type" : "string"},
# 			"type" : {"type" : "string"},
# 			"code" : {
# 				"rgba" : {"type" : "array"},
# 				"hex" : {"type" : "string"}
# 			},
# 			"required" : ["color", "category", "type", "code"]
# 		}
# 	},
# 	"required" : ["colors"]
# }



def validate_color(payload):
        try:
            validate(payload, schema)
            return True
        except Exception as e:
					# print(e)
					return False

q = {
    "color": [
		    {
		      "colr": "black",
		      "category": "hue",
		      "type": "primary",
		      "code": {
		        "rgba": [255,255,255,1],
		        "hex": "#000"
		      }
		    }
    ]
}

s = {
		  "colors": [
		    {
		      "color": "black",
		      "category": "hue",
		      "type": "primary",
		      "code": {
		        "rgba": [255,255,255,1],
		        "hex": "#000"
		      }
		    },
		    {
		      "color": "white",
		      "category": "value",
		      "code": {
		        "rgba": [0,0,0,1],
		        "hex": "#FFF"
		      }
		    },
		    {
		      "color": "red",
		      "category": "hue",
		      "type": "primary",
		      "code": {
		        "rgba": [255,0,0,1],
		        "hex": "#FF0"
		      }
		    }
		  ]
		}

print(validate_color(s))
print(validate_color(q))
