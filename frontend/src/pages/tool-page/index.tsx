import { SidebarProvider, useSidebar } from '@/components/ui/sidebar';
import * as React from 'react';
import { ToolList } from './tool-list';
import { cn } from '@/lib/utils';

const Main = ({ className }: React.ComponentProps<'main'>) => {
  const { open } = useSidebar();
  return <main className={cn('bg-background text-foreground', className)}>Tool Page Content</main>;
};

export const ToolPage: React.FC<{}> = () => {
  return (
    <SidebarProvider className="h-full flex bg-background text-foreground">
      <ToolList className="" />
      <Main className="flex-1 min-h-0 p-4" />
    </SidebarProvider>
  );
};
