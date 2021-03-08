import React, { useEffect } from "react";
import axios from "axios";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
} from "recharts";

// todo: remove hard code
const FETCH_URL = "https://ajf";

interface WakeupData {
  date: string;
  wakeupTime: number;
}

const YAxisMinTime = 4;
const YAxisMaxTime = 10;

// DEMO
const data: WakeupData[] = [
  {
    date: "1/21",
    wakeupTime: 7.5,
  },
  {
    date: "1/22",
    wakeupTime: 8.7,
  },
  {
    date: "1/23",
    wakeupTime: 7.2,
  },
];

/**
 * fetch WakeupData from server.
 * @param user_id
 * @returns  WakeupData[]
 */
const fetchWakeupData = async (user_id: string): Promise<WakeupData[]> => {
  // let results: WakeupData[]
  let results: WakeupData[] | null = await axios
    .get<WakeupData[]>(FETCH_URL, {
      params: {
        user_id: user_id,
      },
    })
    .then<WakeupData[] | null>((res) => {
      return res.data;
    })
    .catch((err) => {
      console.log(err);
      return null;
    });
  if (results === null) {
    throw new Error("can't find WakeupData");
  }
  return results as WakeupData[];
};

export const Dashboard = () => {
  // todo: authから取る
  const user_id = "test";
  useEffect(() => {
    fetchWakeupData(user_id);
  });

  return (
    <>
      <LineChart
        width={500}
        height={300}
        data={data}
        margin={{
          top: 5,
          right: 30,
          left: 20,
          bottom: 5,
        }}
      >
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="date" />
        <YAxis domain={[YAxisMinTime, YAxisMaxTime]} />
        <Tooltip />
        <Legend />
        <Line
          type="monotone"
          dataKey="wakeupTime"
          stroke="#8884d8"
          activeDot={{ r: 8 }}
        />
      </LineChart>
    </>
  );
};
