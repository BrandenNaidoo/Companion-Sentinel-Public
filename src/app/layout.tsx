import type { Metadata } from "next";
import { Inter } from "next/font/google"; // Corrected to use Google Fonts directly
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
    title: "The Companion Sentinel - Server Security for Solopreneurs",
    description: "Automated security checkpoints, vulnerability scanning, and intrusion detection for your VPS fleet.",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en">
            <body className={inter.className}>{children}</body>
        </html>
    );
}
