import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = ((event) => {
    throw error(404, 'Not Found');
}) satisfies PageServerLoad;