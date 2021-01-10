#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

struct Space {
	int w;
	int h;
	int *space;
};

void display(struct Space *space);
void sim(struct Space *);

int main()
{
	int width = 30, height = 30;
	int* screen = (int*)calloc(width*height, sizeof(int));
	struct Space space = {width, height, screen};
	sim(&space);
	return 0;
}

void display(struct Space *space)
{
	printf("\n");
	for (int x = 0; x < space->w; x++) {
		printf("|");
		char pixel;
		for (int y = 0; y < space->h; y++) {
			switch(space->space[(x*space->w)+y]) {
				case 1:
					pixel = 's';
					break;
				case 2:
					pixel = 'w';
					break;
				default:
					pixel = ' ';
					break;
			}
			printf("%c", pixel);
		}
		printf("|\n");
	}
	for (int i = 0; i < (space->w)+2; i++) {
		printf("-");
	}
	return;
}

void sim(struct Space *s)
{
	int width = s->w, height = s->h;
	while(1) {
		for (int i = 0; i < width*height; i++) {
			if (1 == s->space[i]) {
				if (s->space[i+width] != 1) {
					s->space[i] = 0;
					s->space[i+width] = 1;
				} else if (s->space[i+width-1] != 1) {
					s->space[i] = 0;
					s->space[i+width-1] = 1;
				} else if (s->space[i+width+1] != 1) {
					s->space[i] = 0;
					s->space[i+width+1] = 1;
				}
			}
		}

		if (s->space[width/2 + width] != 1) s->space[width/2] = 1;

		display(s);
		sleep(1);
	}
	return;
}
