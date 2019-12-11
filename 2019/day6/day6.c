#include <errno.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/*bool parent(map, parent) {

for k,v in map
if p in v -> k
	else false

}*/
/*
int orbits(map) {
	
	int direct = 0;
	int indirect = 0;

for k, v in map->
	direct += len(v)
	for e in v->
		p = parent(map, k)->
			while p:
				indirect += 1
				p = parent(map, p)
	return direct + indirect;
	
}*/

struct Map {
	char key[4];
	char *data[2048];
	int free_ptr;
};

int main() {
	FILE *fp;
	if ((fp = fopen("test.txt", "r")) == NULL) {
		printf("Couldn't open input: %d\n", errno);
		exit(EXIT_FAILURE);
	}

	struct Map *foo[4096];

	int i = 0;
	char line[12]; // input example: 21X)BWV
	while (fgets(line, sizeof line, fp) != NULL) {
		// remove newline char
		line[strlen(line) - 1] = '\0';

		// split 21X)BWV -> key: 21X, value:BWV
		char *p = NULL, *key = NULL, *value = NULL;
		
		p = strtok(line, ")");
		if (p) key = p;
		p = strtok(NULL, ")");
		if (p) value = p;

		bool y = true;
		for (int i = 0; i < 4096; i++) {
			printf("%s %s\n", key, foo[i]->key);
			if (key == foo[i]->key) {
				y = false;
				break;
			}

		}
		if (y) {
			int free_ptr = 0;
			foo[i] = (struct Map *) malloc(sizeof(struct Map));
			foo[i]->data[free_ptr] = malloc(4096 * sizeof(foo[i]->data[i]));

			foo[i]->free_ptr = free_ptr;
			strcpy(foo[i]->key, key);
			strcpy(foo[i]->data[free_ptr], value);
			printf("%s \n", foo[i]->key);
		} else {
			foo[i]->free_ptr++;
			int free_ptr = foo[i]->free_ptr;
			printf("freeptr %d\n", free_ptr);
			strcpy(foo[i]->data[free_ptr], value);
		}

		i++;
	}
			
	fclose(fp);

	printf("%s %s\n", foo[4]->key, foo[4]->data[1]);
	free(*foo);

	//printf("checksum: %d", orbits(data));
	
	exit(EXIT_SUCCESS);
}
