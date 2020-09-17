CC=gcc
CFLAGS=-O3 -Wall
LIBS=

CFILES=main.c
HFILES=
OFILES=main.o
all:		Lunar

Lunar: 	$(OFILES) $(HFILES)
		$(CC) $(LDFLAGS) $(OFILES) $(LIBS) -o lunar-lander

.c.o:
		$(CC) -c  -I. $(CFLAGS) $(OPTIONS) $<
clean:
		rm -f lunar-lander $(OFILES)
install:
		cp lunar-lander /usr/bin/lunar-lander
sandwich:
		