
 /* Changes made to this file to demonstrate network namespace
 * isolation with containers for a purely educational purpose.
 * Author:
 *   Arun Sriraman <sarun87[at]gmail[dot]com>
 * Original file used from:
 * https://www.redhat.com/archives/libvir-list/2008-January/msg00444.html
 * This file also inherits the license from the original file.
 * Find the original copyright header below
 */

 /*
 * Copyright IBM Corp. 2008
 *
 * lxc_exec.c: example on creating a linux container
 *
 * Authors:
 * David L. Leskovec <dlesko at linux.vnet.ibm.com>
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA 02111-1307  USA
 *
 */

/* System includes */
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>
#include <linux/sched.h>
#include <errno.h>
#include <signal.h>
#include <string.h>

/* Defines */
#define true 1
#define false 0

/* Functions */
static void child_function()
{
    int res;

    /* mount /proc */
    res = mount("lxcproc", "/proc", "proc", 0, NULL);
    if(0 != res)
    {
        printf("Child: mount failed with rc = %d and errno = %d\n", res, errno);
        exit(1);
    }
    else
    {
        printf("Child: mount successful\n");
    }
    system("ps -aef > /tmp/eg_container_ps.out");
    system("ip a > /tmp/eg_container_ipa.out");
}

int main(int argc, char **argv)
{
    int cpid;
    void *childstack, *stack;
    long unsigned int flags;
    int stacksize = getpagesize() * 4;
    int use_newnetns = false;

    if (argc == 2 && strcmp("net",argv[1]) == 0)
        use_newnetns = true;

    /* allocate a stack for the container */
    stack = malloc(stacksize);
    if (!stack)
    {
        printf("Parent: malloc() failed, %s\n", strerror(errno));
        exit(1);
    }
    /* point to "top" of stack */
    childstack = stack + stacksize;

    flags = CLONE_NEWPID|CLONE_NEWNS;
    if (use_newnetns == true)
       flags = flags | CLONE_NEWNET;

    printf("Parent: Clone() flags %lx, pagesize %d...\n",
            flags, getpagesize());

    cpid = clone(child_function, childstack, flags, (void *)argv);
    printf("Parent: cpid: %d\n", cpid);
    if (cpid < 0)
    {
        printf("Parent: clone() failed, %s\n",
                        strerror(errno));
        exit(1);
    }
    system("ip a > /tmp/eg_parent_ipa.out");
    printf("Parent: sleeping, 5 seconds\n");
    sleep(5);
    return 0;
}