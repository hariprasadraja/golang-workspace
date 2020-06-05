1  #pragma once
  2
  3  /*
  4     Copyright (c) 2003 Simon Cooke, All Rights Reserved
  5
  6     Licensed royalty-free for commercial and non-commercial
  7     use, without warranty or guarantee of suitability for any purpose.
  8     All that I ask is that you send me an email
  9     telling me that you're using my code. It'll make me
 10     feel warm and fuzzy inside. spectecjr@gmail.com
 11
 12  */
 13
 14  class BipBuffer
 15  {
 16  private:
 17      BYTE* pBuffer;
 18      int ixa;
 19      int sza;
 20      int ixb;
 21      int szb;
 22      int buflen;
 23      int ixResrv;
 24      int szResrv;
 25
 26  public:
 27      BipBuffer() : pBuffer(NULL), ixa(0), sza(0), ixb(0), szb(0), buflen(0), ixResrv(0), szResrv(0)
 28      {
 29      }
 30
 31      ~BipBuffer()
 32      {
 33          // We don't call FreeBuffer, because we don't need to reset our variables - our object is dying
 34          if (pBuffer != NULL)
 35          {
 36              ::VirtualFree(pBuffer, buflen, MEM_DECOMMIT);
 37          }
 38      }
 39
 40
 41      // Allocate Buffer
 42      //
 43      // Allocates a buffer in virtual memory, to the nearest page size (rounded up)
 44      //
 45      // Parameters:
 46      //   int buffersize                size of buffer to allocate, in bytes (default: 4096)
 47      //
 48      // Returns:
 49      //   bool                        true if successful, false if buffer cannot be allocated
 50
 51      bool AllocateBuffer(int buffersize = 4096)
 52      {
 53          if (buffersize <= 0) return false;
 54
 55          if (pBuffer != NULL) FreeBuffer();
 56
 57          SYSTEM_INFO si;
 58          ::GetSystemInfo(&si);
 59
 60          // Calculate nearest page size
 61          buffersize = ((buffersize + si.dwPageSize - 1) / si.dwPageSize) * si.dwPageSize;
 62
 63          pBuffer = (BYTE*)::VirtualAlloc(NULL, buffersize, MEM_COMMIT, PAGE_READWRITE);
 64          if (pBuffer == NULL) return false;
 65
 66          buflen = buffersize;
 67          return true;
 68      }


 69
 70      ///
 71      /// \brief Clears the buffer of any allocations.
 72      ///
 73      /// Clears the buffer of any allocations or reservations. Note; it
 74      /// does not wipe the buffer memory; it merely resets all pointers,
 75      /// returning the buffer to a completely empty state ready for new
 76      /// allocations.
 77      ///
 78
 79      void Clear()
 80      {
 81          ixa = sza = ixb = szb = ixResrv = szResrv = 0;
 82      }
 83
 84      // Free Buffer
 85      //
 86      // Frees a previously allocated buffer, resetting all internal pointers to 0.
 87      //
 88      // Parameters:
 89      //   none
 90      //
 91      // Returns:
 92      //   void
 93
 94      void FreeBuffer()
 95      {
 96          if (pBuffer == NULL) return;
 97
 98          ixa = sza = ixb = szb = buflen = 0;
 99
100          ::VirtualFree(pBuffer, buflen, MEM_DECOMMIT);
101
102          pBuffer = NULL;
103      }
104
105      // Reserve
106      //
107      // Reserves space in the buffer for a memory write operation
108      //
109      // Parameters:
110      //   int size                amount of space to reserve
111      //   OUT int& reserved        size of space actually reserved
112      //
113      // Returns:
114      //   BYTE*                    pointer to the reserved block
115      //
116      // Notes:
117      //   Will return NULL for the pointer if no space can be allocated.
118      //   Can return any value from 1 to size in reserved.
119      //   Will return NULL if a previous reservation has not been committed.
120
121      BYTE* Reserve(int size, OUT int& reserved)
122      {
123          // We always allocate on B if B exists; this means we have two blocks and our buffer is filling.
124          if (szb)
125          {
126              int freespace = GetBFreeSpace();
127
128              if (size < freespace) freespace = size;
129
130              if (freespace == 0) return NULL;
131
132              szResrv = freespace;
133              reserved = freespace;
134              ixResrv = ixb + szb;
135              return pBuffer + ixResrv;
136          }
137          else
138          {
139              // Block b does not exist, so we can check if the space AFTER a is bigger than the space
140              // before A, and allocate the bigger one.
141
142              int freespace = GetSpaceAfterA();
143              if (freespace >= ixa)
144              {
145                  if (freespace == 0) return NULL;
146                  if (size < freespace) freespace = size;
147
148                  szResrv = freespace;
149                  reserved = freespace;
150                  ixResrv = ixa + sza;
151                  return pBuffer + ixResrv;
152              }
153              else
154              {
155                  if (ixa == 0) return NULL;
156                  if (ixa < size) size = ixa;
157                  szResrv = size;
158                  reserved = size;
159                  ixResrv = 0;
160                  return pBuffer;
161              }
162          }
163      }
164
165      // Commit
166      //
167      // Commits space that has been written to in the buffer
168      //
169      // Parameters:
170      //   int size                number of bytes to commit
171      //
172      // Notes:
173      //   Committing a size > than the reserved size will cause an assert in a debug build;
174      //   in a release build, the actual reserved size will be used.
175      //   Committing a size < than the reserved size will commit that amount of data, and release
176      //   the rest of the space.
177      //   Committing a size of 0 will release the reservation.
178      //
179
180      void Commit(int size)
181      {
182          if (size == 0)
183          {
184              // decommit any reservation
185              szResrv = ixResrv = 0;
186              return;
187          }
188
189          // If we try to commit more space than we asked for, clip to the size we asked for.
190
191          if (size > szResrv)
192          {
193              size = szResrv;
194          }
195
196          // If we have no blocks being used currently, we create one in A.
197
198          if (sza == 0 && szb == 0)
199          {
200              ixa = ixResrv;
201              sza = size;
202
203              ixResrv = 0;
204              szResrv = 0;
205              return;
206          }
207
208          if (ixResrv == sza + ixa)
209          {
210              sza += size;
211          }
212          else
213          {
214              szb += size;
215          }
216
217          ixResrv = 0;
218          szResrv = 0;
219      }
220
221      // GetContiguousBlock
222      //
223      // Gets a pointer to the first contiguous block in the buffer, and returns the size of that block.
224      //
225      // Parameters:
226      //   OUT int & size            returns the size of the first contiguous block
227      //
228      // Returns:
229      //   BYTE*                    pointer to the first contiguous block, or NULL if empty.
230
231      BYTE* GetContiguousBlock(OUT int& size)
232      {
233          if (sza == 0)
234          {
235              size = 0;
236              return NULL;
237          }
238
239          size = sza;
240          return pBuffer + ixa;
241
242      }
243
244      // DecommitBlock
245      //
246      // Decommits space from the first contiguous block
247      //
248      // Parameters:
249      //   int size                amount of memory to decommit
250      //
251      // Returns:
252      //   nothing
253
254      void DecommitBlock(int size)
255      {
256          if (size >= sza)
257          {
258              ixa = ixb;
259              sza = szb;
260              ixb = 0;
261              szb = 0;
262          }
263          else
264          {
265              sza -= size;
266              ixa += size;
267          }
268      }
269
270      // GetCommittedSize
271      //
272      // Queries how much data (in total) has been committed in the buffer
273      //
274      // Parameters:
275      //   none
276      //
277      // Returns:
278      //   int                    total amount of committed data in the buffer
279
280      int    GetCommittedSize() const
281      {
282          return sza + szb;
283      }
284
285      // GetReservationSize
286      //
287      // Queries how much space has been reserved in the buffer.
288      //
289      // Parameters:
290      //   none
291      //
292      // Returns:
293      //   int                    number of bytes that have been reserved
294      //
295      // Notes:
296      //   A return value of 0 indicates that no space has been reserved
297
298      int GetReservationSize() const
299      {
300          return szResrv;
301      }
302
303      // GetBufferSize
304      //
305      // Queries the maximum total size of the buffer
306      //
307      // Parameters:
308      //   none
309      //
310      // Returns:
311      //   int                    total size of buffer
312
313      int GetBufferSize() const
314      {
315          return buflen;
316      }
317
318      // IsInitialized
319      //
320      // Queries whether or not the buffer has been allocated
321      //
322      // Parameters:
323      //   none
324      //
325      // Returns:
326      //   bool                    true if the buffer has been allocated
327
328      bool IsInitialized() const
329      {
330          return pBuffer != NULL;
331      }
332
333  private:
334      int GetSpaceAfterA() const
335      {
336          return buflen - ixa - sza;
337      }
338
339      int GetBFreeSpace() const
340      {
341          return ixa - ixb - szb;
342      }
343  };