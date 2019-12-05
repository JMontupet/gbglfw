package glfwio

const vertexShader = `
#version 330

uniform mat4 projection;

in vec3 vert;
in vec2 vertTexCoord;

out vec2 fragTexCoord;

void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * vec4(vert, 1);
}
` + "\x00"

const fragmentShaderScale3x = `
#version 330

uniform sampler2D tex;
uniform sampler1D pal;

in vec2 fragTexCoord;

out vec4 outputColor;

vec2 step = vec2(1.0 / 160, 1.0 / 144);
vec4 getColorIndex(in vec2 coord) {
	vec2 floorCoord;
	floorCoord.x = floor(coord.x * 160) / 160 + step.x/2;
	floorCoord.y = floor(coord.y * 144) / 144 + step.y/2;

	vec2 roundCoord;
	roundCoord.x = floor(coord.x * 160) / 160;
	roundCoord.y = floor(coord.y * 144) / 144;

	vec4 A = texture(tex, vec2(floorCoord.x-step.x, floorCoord.y-step.y));
	vec4 B = texture(tex, vec2(floorCoord.x       , floorCoord.y-step.y));
	vec4 C = texture(tex, vec2(floorCoord.x+step.x, floorCoord.y-step.y));
	vec4 D = texture(tex, vec2(floorCoord.x-step.x, floorCoord.y       ));
	vec4 E = texture(tex, vec2(floorCoord.x       , floorCoord.y       ));
	vec4 F = texture(tex, vec2(floorCoord.x+step.x, floorCoord.y       ));
	vec4 G = texture(tex, vec2(floorCoord.x-step.x, floorCoord.y+step.y));
	vec4 H = texture(tex, vec2(floorCoord.x       , floorCoord.y+step.y));
	vec4 I = texture(tex, vec2(floorCoord.x+step.x, floorCoord.y+step.y));

	vec2 step3 = step;
	step3.x = step.x / 3;
	step3.y = step.y / 3;

	// Line 1
	if (coord.y < roundCoord.y+step3.y) {
		// P1
		if (coord.x < roundCoord.x+step3.x) {
			if (D==B && D!=H && B!=F) {
				return D;
			}
			return E;

		}
		// P2
		else if(coord.x < roundCoord.x+step3.x*2) {
			if ((D==B && D!=H && B!=F && E!=C) || (B==F && B!=D && F!=H && E!=A)) {
				return B;
			}
			return E;

		}
		// P3
		else {
			if (B==F && B!=D && F!=H) {
				return F;
			}
			return E;

		}
	}
	// Line 2
	else if (coord.y < roundCoord.y+step3.y*2) {
		// P4
		if (coord.x < roundCoord.x+step3.x) {
			if ((H==D && H!=F && D!=B && E!=A) || (D==B && D!=H && B!=F && E!=G)) {
				return D;
			}
			return E;

		}
		// P5
		else if(coord.x < roundCoord.x+step3.x*2) {
			return E;
		}
		// P6
		else {
			if ((B==F && B!=D && F!=H && E!=I) || (F==H && F!=B && H!=D && E!=C)) {
				return F;
			}
			return E;

		}
	}
	// Line 3
	else  {

		// P7
		if (coord.x < roundCoord.x+step3.x) {
			if (H==D && H!=F && D!=B) {
				return D;
			}
			return E;

		}
		// P8
		else if(coord.x < roundCoord.x+step3.x*2) {
			if ((F==H && F!=B && H!=D && E!=G) || (H==D && H!=F && D!=B && E!=I)) {
				return H;
			}
			return E;

		}
		// P9
		else {

			if (F==H && F!=B && H!=D) {

				return F;
			}
			return E;

		}
	} 
	return vec4(1.0,0.0,0.0, 1.0);
	
	return E;
}
const mat3 GBCMatrix = mat3( 0.924, 0.021, 0.013, 0.048, 0.787, 0.249, 0.104, 0.09, 0.733 );
const float gamma = 2.2;
void main() {
	vec4 index = getColorIndex(fragTexCoord);
	outputColor = texture(pal, index.r * 4.0);

	outputColor.rgb = pow(outputColor.rgb, vec3(gamma));
	vec3 Picture = outputColor.xyz;
    // Picture *= Picture;
    Picture *= GBCMatrix;
    // Picture = sqrt(Picture);
	outputColor = vec4(Picture,1.0);
	outputColor.rgb = pow(outputColor.rgb, vec3(1/gamma));
}
` + "\x00"

const fragmentShaderNearest = `
#version 330

uniform sampler2D tex;
uniform sampler1D pal;

in vec2 fragTexCoord;

out vec4 outputColor;

// Magic matrix
const mat3 GBCMatrix = mat3( 0.924, 0.021, 0.013, 0.048, 0.787, 0.249, 0.104, 0.09, 0.733 );
const float gamma = 2.2;

void main() {
	// Apply Color Palette
	vec4 index = texture(tex, fragTexCoord);
	outputColor = texture(pal, index.r * 4.0 + (1.0/64.0/2.0));

	// Color Correction
	outputColor.rgb = pow(outputColor.rgb, vec3(gamma));
	vec3 Picture = outputColor.xyz;
    // Picture *= Picture;
    Picture *= GBCMatrix;
    // Picture = sqrt(Picture);
	outputColor = vec4(Picture,1.0);
	outputColor.rgb = pow(outputColor.rgb, vec3(1/gamma));
	

}
` + "\x00"
