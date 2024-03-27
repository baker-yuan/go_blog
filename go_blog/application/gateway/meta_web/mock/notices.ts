import { Request, Response } from 'express';

const getNotices = (req: Request, res: Response) => {
  res.json([]);
};

export default {
  'GET /api/notices': getNotices,
};
