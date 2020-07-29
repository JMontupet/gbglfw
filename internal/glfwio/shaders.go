package glfwio

const vertexShader = `
#version 110

uniform mat4 projection;

attribute vec3 vert;
attribute vec2 vertTexCoord;

varying vec2 fragTexCoord;

void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * vec4(vert, 1);
}
` + "\x00"

const fragmentShaderNearest = `
#version 110

uniform sampler2D tex;
uniform sampler1D pal;

varying vec2 fragTexCoord;


// Magic matrix
const mat3 GBCMatrix = mat3( 0.924, 0.021, 0.013, 0.048, 0.787, 0.249, 0.104, 0.09, 0.733 );
const float gamma = 2.2;

void main() {
	vec4 outputColor;
	// Apply Color Palette
	vec4 index = texture2D(tex, fragTexCoord);
	outputColor = texture1D(pal, index.r * 4.0 + (1.0/64.0/2.0));

	// Color Correction
	outputColor.rgb = pow(outputColor.rgb, vec3(gamma));
	vec3 Picture = outputColor.xyz;
    // Picture *= Picture;
    Picture *= GBCMatrix;
    // Picture = sqrt(Picture);
	outputColor = vec4(Picture,1.0);
	outputColor.rgb = pow(outputColor.rgb, vec3(1.0/gamma));
	gl_FragColor = outputColor;
}
` + "\x00"
